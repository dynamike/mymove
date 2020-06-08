package paymentrequest

import (
	"fmt"
	"io"
	"path"
	"time"

	"github.com/transcom/mymove/pkg/services"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/storage"
	"github.com/transcom/mymove/pkg/uploader"
)

type paymentRequestUploadCreator struct {
	db            *pop.Connection
	logger        Logger
	fileStorer    storage.FileStorer
	fileSizeLimit uploader.ByteSize
}

// NewPaymentRequestUploadCreator returns a new payment request upload creator
func NewPaymentRequestUploadCreator(db *pop.Connection, logger Logger, fileStorer storage.FileStorer) services.PaymentRequestUploadCreator {
	return &paymentRequestUploadCreator{db, logger, fileStorer, uploader.MaxFileSizeLimit}
}

func (p *paymentRequestUploadCreator) assembleUploadFilePathName(paymentRequestID uuid.UUID) (string, error) {
	var paymentRequest models.PaymentRequest
	err := p.db.Where("id=$1", paymentRequestID).First(&paymentRequest)
	if err != nil {
		return "", fmt.Errorf("cannot fetch payment request: %w", err)
	}

	filename := "timestamp-" + time.Now().String()
	uploadFilePath := fmt.Sprintf("/app/payment-request-uploads/mto-%s/payment-request-%s", paymentRequest.MoveTaskOrderID, paymentRequest.ID)
	uploadFileName := path.Join(uploadFilePath, filename)

	return uploadFileName, err
}

func (p *paymentRequestUploadCreator) CreateUpload(file io.ReadCloser, paymentRequestID uuid.UUID, contractorID uuid.UUID) (*models.Upload, error) {
	var upload *models.Upload
	transactionError := p.db.Transaction(func(tx *pop.Connection) error {
		newUploader, err := uploader.NewPrimeUploader(tx, p.logger, p.fileStorer, p.fileSizeLimit)
		if err != nil {
			return fmt.Errorf("cannot create uploader in payment request uploadCreator: %w", err)
		}

		fileName, err := p.assembleUploadFilePathName(paymentRequestID)
		if err != nil {
			return fmt.Errorf("could not assemble primeUpload filepath name %w", err)
		}

		aFile, err := newUploader.PrepareFileForUpload(file, fileName)
		if err != nil {
			return fmt.Errorf("could not prepare file for uploader: %w", err)
		}

		newUploader.SetUploadStorageKey(fileName)

		var paymentRequest models.PaymentRequest
		err = tx.Find(&paymentRequest, paymentRequestID)
		if err != nil {
			return fmt.Errorf("could not find PaymentRequestID [%s]: %w", paymentRequestID, err)
		}
		// create proof of service doc
		proofOfServiceDoc := models.ProofOfServiceDoc{
			PaymentRequestID: paymentRequestID,
			PaymentRequest:   paymentRequest,
		}
		verrs, err := tx.ValidateAndCreate(&proofOfServiceDoc)
		if err != nil {
			return fmt.Errorf("failure creating proof of service doc: %w", err)
		}
		if verrs.HasAny() {
			return fmt.Errorf("validation error creating proof of service doc: %w", verrs)
		}

		posID := &proofOfServiceDoc.ID
		primeUpload, verrs, err := newUploader.CreatePrimeUploadForDocument(posID, contractorID, uploader.File{File: aFile}, uploader.AllowedTypesPaymentRequest)
		if verrs.HasAny() {
			return fmt.Errorf("validation error creating payment request primeUpload: %w", verrs)
		}
		if err != nil {
			return fmt.Errorf("failure creating payment request primeUpload: %w", err)
		}
		upload = &primeUpload.Upload
		return nil
	})

	if transactionError != nil {
		return nil, transactionError
	}

	return upload, nil
}
