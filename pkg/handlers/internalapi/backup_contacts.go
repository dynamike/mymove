package internalapi

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"

	backupop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/backup_contacts"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
)

func payloadForBackupContactModel(contact models.BackupContact) internalmessages.ServiceMemberBackupContactPayload {
	permission := internalmessages.NewBackupContactPermission(internalmessages.BackupContactPermission(contact.Permission))
	contactPayload := internalmessages.ServiceMemberBackupContactPayload{
		ID:              handlers.FmtUUID(contact.ID),
		ServiceMemberID: *handlers.FmtUUID(contact.ServiceMemberID),
		UpdatedAt:       handlers.FmtDateTime(contact.UpdatedAt),
		CreatedAt:       handlers.FmtDateTime(contact.CreatedAt),
		Name:            &contact.Name,
		Email:           &contact.Email,
		Telephone:       contact.Phone,
		Permission:      permission,
	}
	return contactPayload
}

// CreateBackupContactHandler creates a new backup contact
type CreateBackupContactHandler struct {
	handlers.HandlerContext
}

// Handle ... creates a new BackupContact from a request payload
func (h CreateBackupContactHandler) Handle(params backupop.CreateServiceMemberBackupContactParams) middleware.Responder {

	ctx := params.HTTPRequest.Context()

	// User should always be populated by middleware
	session, logger := h.SessionAndLoggerFromContext(ctx)
	serviceMemberID, _ := uuid.FromString(params.ServiceMemberID.String())
	serviceMember, err := models.FetchServiceMemberForUser(h.DB(), session, serviceMemberID)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	if params.CreateBackupContactPayload.Permission == nil {
		return handlers.ResponseForError(logger, errors.New("missing required field: Permission"))
	}

	newContact, verrs, err := serviceMember.CreateBackupContact(h.DB(),
		*params.CreateBackupContactPayload.Name,
		*params.CreateBackupContactPayload.Email,
		params.CreateBackupContactPayload.Telephone,
		models.BackupContactPermission(*params.CreateBackupContactPayload.Permission))
	if err != nil || verrs.HasAny() {
		return handlers.ResponseForVErrors(logger, verrs, err)
	}

	contactPayload := payloadForBackupContactModel(newContact)
	return backupop.NewCreateServiceMemberBackupContactCreated().WithPayload(&contactPayload)
}

// IndexBackupContactsHandler returns a list of all backup contacts for a service member
type IndexBackupContactsHandler struct {
	handlers.HandlerContext
}

// Handle retrieves a list of all moves in the system belonging to the logged in user
func (h IndexBackupContactsHandler) Handle(params backupop.IndexServiceMemberBackupContactsParams) middleware.Responder {
	session, logger := h.SessionAndLoggerFromRequest(params.HTTPRequest)

	serviceMemberID, _ := uuid.FromString(params.ServiceMemberID.String())
	serviceMember, err := models.FetchServiceMemberForUser(h.DB(), session, serviceMemberID)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	contacts := serviceMember.BackupContacts

	contactPayloads := make(internalmessages.IndexServiceMemberBackupContactsPayload, len(contacts))
	for i, contact := range contacts {
		contactPayload := payloadForBackupContactModel(contact)
		contactPayloads[i] = &contactPayload
	}

	return backupop.NewIndexServiceMemberBackupContactsOK().WithPayload(contactPayloads)
}

// ShowBackupContactHandler returns a backup contact for a user and backup contact ID
type ShowBackupContactHandler struct {
	handlers.HandlerContext
}

// Handle retrieves a backup contact in the system belonging to the logged in user given backup contact ID
func (h ShowBackupContactHandler) Handle(params backupop.ShowServiceMemberBackupContactParams) middleware.Responder {
	// User should always be populated by middleware
	session, logger := h.SessionAndLoggerFromRequest(params.HTTPRequest)
	contactID, _ := uuid.FromString(params.BackupContactID.String())
	contact, err := models.FetchBackupContact(h.DB(), session, contactID)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	contactPayload := payloadForBackupContactModel(contact)
	return backupop.NewShowServiceMemberBackupContactOK().WithPayload(&contactPayload)
}

// UpdateBackupContactHandler updates a backup contact with a new one
type UpdateBackupContactHandler struct {
	handlers.HandlerContext
}

// Handle ... updates a BackupContact from a request payload
func (h UpdateBackupContactHandler) Handle(params backupop.UpdateServiceMemberBackupContactParams) middleware.Responder {
	// User should always be populated by middleware
	session, logger := h.SessionAndLoggerFromRequest(params.HTTPRequest)
	contactID, _ := uuid.FromString(params.BackupContactID.String())
	contact, err := models.FetchBackupContact(h.DB(), session, contactID)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	contact.Name = *params.UpdateServiceMemberBackupContactPayload.Name
	contact.Email = *params.UpdateServiceMemberBackupContactPayload.Email
	contact.Phone = params.UpdateServiceMemberBackupContactPayload.Telephone
	if params.UpdateServiceMemberBackupContactPayload.Permission != nil {
		contact.Permission = models.BackupContactPermission(*params.UpdateServiceMemberBackupContactPayload.Permission)
	}

	if verrs, err := h.DB().ValidateAndUpdate(&contact); verrs.HasAny() || err != nil {
		return handlers.ResponseForVErrors(logger, verrs, err)
	}

	contactPayload := payloadForBackupContactModel(contact)
	return backupop.NewUpdateServiceMemberBackupContactCreated().WithPayload(&contactPayload)
}
