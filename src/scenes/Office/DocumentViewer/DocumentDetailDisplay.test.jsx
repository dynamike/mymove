import { shallow } from 'enzyme';
import DocumentDetailDisplay from './DocumentDetailDisplay';
import React from 'react';
import { MOVE_DOC_TYPE, WEIGHT_TICKET_SET_TYPE } from '../../../shared/constants';

describe('DocumentDetailDisplay', () => {
  const renderDocumentDetailDisplay = ({
    isExpenseDocument = false,
    isWeightTicketDocument = true,
    moveDocument = {},
    moveDocSchema = {},
    isStorageExpenseDocument = false,
  }) =>
    shallow(
      <DocumentDetailDisplay
        isExpenseDocument={isExpenseDocument}
        isWeightTicketDocument={isWeightTicketDocument}
        moveDocument={moveDocument}
        moveDocSchema={moveDocSchema}
        isStorageExpenseDocument={isStorageExpenseDocument}
      />,
    );

  describe('weight ticket document display view', () => {
    const requiredMoveDocumentFields = {
      id: 'id',
      move_id: 'id',
      move_document_type: MOVE_DOC_TYPE.WEIGHT_TICKET_SET,
      document: {
        id: 'an id',
        move_document_id: 'another id',
        service_member_id: 'another id2',
        uploads: [
          {
            id: 'id',
            url: 'url',
            filename: 'file here',
            content_type: 'json',
            created_at: '2018-09-27 08:14:38.702434',
          },
        ],
      },
    };
    it('includes common document info', () => {
      const moveDocument = Object.assign(requiredMoveDocumentFields, {
        title: 'My Title',
        notes: 'This is a note',
        status: 'AWAITING_REVIEW',
      });

      const moveDocSchema = {
        properties: {
          title: { enum: false },
          move_document_type: { enum: false },
          status: { enum: false },
          notes: { enum: false },
        },
        required: [],
        type: 'string type',
      };

      const documentDisplay = renderDocumentDetailDisplay({ moveDocument, moveDocSchema });
      expect(documentDisplay.find('[data-cy="panel-subhead"]').text()).toContain(moveDocument.title);
      expect(documentDisplay.find('[data-cy="uploaded-at"]').text()).toEqual('Uploaded 27-Sep-18');
      expect(
        documentDisplay.find('[data-cy="document-title"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.title);
      expect(
        documentDisplay.find('[data-cy="move-document-type"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.move_document_type);
      expect(documentDisplay.find('[data-cy="status"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.status,
      );
      expect(documentDisplay.find('[data-cy="notes"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.notes,
      );
    });

    it('includes weight ticket-specific fields', () => {
      const documentFieldsToTest = {
        empty_weight: '2200',
        full_weight: '3500',
      };

      const moveDocSchema = {
        properties: {
          empty_weight: { enum: false },
          full_weight: { enum: false },
        },
        required: [],
        type: 'string type',
      };

      const moveDocument = Object.assign(requiredMoveDocumentFields, documentFieldsToTest);
      const documentDisplay = renderDocumentDetailDisplay({ moveDocument, moveDocSchema });

      expect(documentDisplay.find('[data-cy="empty-weight"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.empty_weight,
      );
      expect(documentDisplay.find('[data-cy="full-weight"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.full_weight,
      );
    });

    describe('is car or car and trailer', () => {
      it('includes the make and model fields ', () => {
        const documentFieldsToTest = {
          vehicle_make: 'Honda',
          vehicle_model: 'Civic',
          weight_ticket_set_type: WEIGHT_TICKET_SET_TYPE.CAR,
        };

        const moveDocSchema = {
          properties: {
            weight_ticket_set_type: { enum: false },
            vehicle_make: { enum: false },
            vehicle_model: { enum: false },
          },
          required: [],
          type: 'string type',
        };

        const moveDocument = Object.assign(requiredMoveDocumentFields, documentFieldsToTest);
        const documentDisplay = renderDocumentDetailDisplay({ moveDocument, moveDocSchema });
        expect(
          documentDisplay.find('[data-cy="weight-ticket-set-type"]').dive().dive().find('SwaggerValue').dive().text(),
        ).toEqual(moveDocument.weight_ticket_set_type);
        expect(
          documentDisplay.find('[data-cy="vehicle-make"]').dive().dive().find('SwaggerValue').dive().text(),
        ).toEqual(moveDocument.vehicle_make);
        expect(
          documentDisplay.find('[data-cy="vehicle-model"]').dive().dive().find('SwaggerValue').dive().text(),
        ).toEqual(moveDocument.vehicle_model);
      });
    });

    describe('a box truck type weight ticket', () => {
      it('includes vehicle nickname', () => {
        const documentFieldsToTest = {
          vehicle_nickname: '15 foot box truck',
          weight_ticket_set_type: WEIGHT_TICKET_SET_TYPE.BOX_TRUCK,
        };

        const moveDocSchema = {
          properties: {
            weight_ticket_set_type: { enum: false },
            vehicle_nickname: { enum: false },
          },
          required: [],
          type: 'string type',
        };

        const moveDocument = Object.assign(requiredMoveDocumentFields, documentFieldsToTest);
        const documentDisplay = renderDocumentDetailDisplay({ moveDocument, moveDocSchema });
        expect(
          documentDisplay.find('[data-cy="weight-ticket-set-type"]').dive().dive().find('SwaggerValue').dive().text(),
        ).toEqual(moveDocument.weight_ticket_set_type);
        expect(
          documentDisplay.find('[data-cy="vehicle-nickname"]').dive().dive().find('SwaggerValue').dive().text(),
        ).toEqual(moveDocument.vehicle_nickname);
      });
    });
  });
  describe('expense document display view', () => {
    const requiredMoveDocumentFields = {
      id: 'id',
      move_id: 'id',
      move_document_type: MOVE_DOC_TYPE.EXPENSE,
      document: {
        id: 'an id',
        move_document_id: 'another id',
        service_member_id: 'another id2',
        uploads: [
          {
            id: 'id',
            url: 'url',
            filename: 'file here',
            content_type: 'json',
            created_at: '2018-09-27 08:14:38.702434',
          },
        ],
      },
    };
    it('has all expected fields', () => {
      const moveDocument = Object.assign(requiredMoveDocumentFields, {
        title: 'My Title',
        move_document_type: MOVE_DOC_TYPE.EXPENSE,
        moving_expense_type: 'RENTAL_EQUIPMENT',
        requested_amount_cents: '45000',
        payment_method: 'GCCC',
        notes: 'This is a note',
        status: 'AWAITING_REVIEW',
      });

      const moveDocSchema = {
        properties: {
          title: { enum: false },
          move_document_type: { enum: false },
          moving_expense_type: { enum: false },
          requested_amount_cents: { enum: false },
          payment_method: { enum: false },
          status: { enum: false },
          notes: { enum: false },
        },
        required: [],
        type: 'string type',
      };

      const documentDisplay = renderDocumentDetailDisplay({
        isExpenseDocument: true,
        isWeightTicketDocument: false,
        moveDocument,
        moveDocSchema,
      });
      expect(
        documentDisplay.find('[data-cy="document-title"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.title);
      expect(
        documentDisplay.find('[data-cy="move-document-type"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.move_document_type);
      expect(
        documentDisplay.find('[data-cy="moving-expense-type"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.moving_expense_type);
      expect(
        documentDisplay.find('[data-cy="requested-amount-cents"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.requested_amount_cents);
      expect(
        documentDisplay.find('[data-cy="payment-method"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.payment_method);
      expect(documentDisplay.find('[data-cy="status"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.status,
      );
      expect(documentDisplay.find('[data-cy="notes"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.notes,
      );
    });
  });
  describe('storage expense document display view', () => {
    const requiredMoveDocumentFields = {
      id: 'id',
      move_id: 'id',
      move_document_type: MOVE_DOC_TYPE.EXPENSE,
      document: {
        id: 'an id',
        move_document_id: 'another id',
        service_member_id: 'another id2',
        uploads: [
          {
            id: 'id',
            url: 'url',
            filename: 'file here',
            content_type: 'json',
            created_at: '2018-09-27 08:14:38.702434',
          },
        ],
      },
    };
    it('has all expected fields', () => {
      const moveDocument = Object.assign(requiredMoveDocumentFields, {
        title: 'My Title',
        move_document_type: 'STORAGE_EXPENSE',
        storage_start_date: '2018-09-27 08:14:38.702434',
        storage_end_date: '2018-10-25 08:14:38.702434',
        notes: 'This is a note',
        status: 'AWAITING_REVIEW',
      });

      const moveDocSchema = {
        properties: {
          title: { enum: false },
          move_document_type: { enum: false },
          storage_start_date: { enum: false },
          storage_end_date: { enum: false },
          status: { enum: false },
          notes: { enum: false },
        },
        required: [],
        type: 'string type',
      };

      const documentDisplay = renderDocumentDetailDisplay({
        isExpenseDocument: false,
        isWeightTicketDocument: false,
        isStorageExpenseDocument: true,
        moveDocument,
        moveDocSchema,
      });
      expect(
        documentDisplay.find('[data-cy="document-title"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.title);
      expect(
        documentDisplay.find('[data-cy="move-document-type"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.move_document_type);
      expect(
        documentDisplay.find('[data-cy="storage-start-date"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.storage_start_date);
      expect(
        documentDisplay.find('[data-cy="storage-end-date"]').dive().dive().find('SwaggerValue').dive().text(),
      ).toEqual(moveDocument.storage_end_date);
      expect(documentDisplay.find('[data-cy="status"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.status,
      );
      expect(documentDisplay.find('[data-cy="notes"]').dive().dive().find('SwaggerValue').dive().text()).toEqual(
        moveDocument.notes,
      );
    });
  });
});
