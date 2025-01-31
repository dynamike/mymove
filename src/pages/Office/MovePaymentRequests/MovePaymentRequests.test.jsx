/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import { render, waitFor, screen, within } from '@testing-library/react';

import MovePaymentRequests from './MovePaymentRequests';

import { MockProviders } from 'testUtils';
import { useMovePaymentRequestsQueries } from 'hooks/queries';
import { shipmentStatuses } from 'constants/shipments';
import SERVICE_ITEM_STATUSES from 'constants/serviceItems';

jest.mock('hooks/queries', () => ({
  useMovePaymentRequestsQueries: jest.fn(),
}));

const testProps = {
  setUnapprovedShipmentCount: jest.fn(),
  setUnapprovedServiceItemCount: jest.fn(),
  setPendingPaymentRequestCount: jest.fn(),
};

const move = {
  id: '1',
  contractor: {
    contractNumber: 'HTC-123-3456',
  },
  orders: {
    sac: '1234456',
    tac: '1213',
  },
};

const multiplePaymentRequests = {
  paymentRequests: [
    {
      id: '09474c6a-69b6-4501-8e08-670a12512e5f',
      createdAt: '2020-12-01T00:00:00.000Z',
      paymentRequestNumber: '1843-9061-1',
      status: 'REVIEWED',
      moveTaskOrderID: '1',
      moveTaskOrder: move,
      serviceItems: [
        {
          id: '09474c6a-69b6-4501-8e08-670a12512a5f',
          createdAt: '2020-12-01T00:00:00.000Z',
          mtoServiceItemID: 'f8c2f97f-99e7-4fb1-9cc4-473debd24dbc',
          priceCents: 2000001,
          status: 'APPROVED',
        },
        {
          id: '39474c6a-69b6-4501-8e08-670a12512a5f',
          createdAt: '2020-12-01T00:00:00.000Z',
          mtoServiceItemID: 'a8c2f97f-99e7-4fb1-9cc4-473debd24dbc',
          priceCents: 4000001,
          status: 'DENIED',
          rejectionReason: 'Requested amount exceeds guideline',
        },
      ],
      reviewedAt: '2020-12-01T00:00:00.000Z',
    },
    {
      id: '29474c6a-69b6-4501-8e08-670a12512e5f',
      createdAt: '2020-12-01T00:00:00.000Z',
      paymentRequestNumber: '1843-9061-2',
      status: 'PENDING',
      moveTaskOrderID: '1',
      moveTaskOrder: move,
      serviceItems: [
        {
          id: '09474c6a-69b6-4501-8e08-670a12512a5f',
          createdAt: '2020-12-01T00:00:00.000Z',
          mtoServiceItemID: 'f8c2f97f-99e7-4fb1-9cc4-473debd24dbc',
          priceCents: 2000001,
          status: 'REQUESTED',
        },
        {
          id: '39474c6a-69b6-4501-8e08-670a12512a5f',
          createdAt: '2020-12-01T00:00:00.000Z',
          mtoServiceItemID: 'a8c2f97f-99e7-4fb1-9cc4-473debd24dbc',
          priceCents: 4000001,
          status: 'REQUESTED',
        },
      ],
    },
  ],
  mtoShipments: [
    {
      id: '2',
      moveTaskOrderID: '1',
      status: shipmentStatuses.APPROVED,
      scheduledPickupDate: '2020-01-09T00:00:00.000Z',
      destinationAddress: { city: 'Princeton', state: 'NJ', postal_code: '08540' },
      pickupAddress: { city: 'Boston', state: 'MA', postal_code: '02101' },
      mtoServiceItems: [
        {
          id: '5',
          mtoShipmentID: '2',
          status: SERVICE_ITEM_STATUSES.APPROVED,
        },
        {
          id: '6',
          status: SERVICE_ITEM_STATUSES.REJECTED,
          mtoShipmentID: '2',
        },
        {
          id: '7',
          status: SERVICE_ITEM_STATUSES.SUBMITTED,
          mtoShipmentID: '2',
        },
      ],
    },
    {
      id: '3',
      moveTaskOrderID: '1',
      status: shipmentStatuses.APPROVED,
      scheduledPickupDate: '2020-01-10T00:00:00.000Z',
      destinationAddress: { city: 'Princeton', state: 'NJ', postal_code: '08540' },
      pickupAddress: { city: 'Boston', state: 'MA', postal_code: '02101' },
      mtoServiceItems: [
        {
          id: '9',
          mtoShipmentID: '3',
          status: SERVICE_ITEM_STATUSES.APPROVED,
        },
        {
          id: '10',
          status: SERVICE_ITEM_STATUSES.REJECTED,
          mtoShipmentID: '3',
        },
        {
          id: '11',
          status: SERVICE_ITEM_STATUSES.SUBMITTED,
          mtoShipmentID: '3',
        },
      ],
    },
    {
      id: '4',
      moveTaskOrderID: '1',
      status: shipmentStatuses.SUBMITTED,
      scheduledPickupDate: '2020-01-11T00:00:00.000Z',
      destinationAddress: { city: 'Princeton', state: 'NJ', postal_code: '08540' },
      pickupAddress: { city: 'Boston', state: 'MA', postal_code: '02101' },
      mtoServiceItems: [
        {
          id: '12',
          mtoShipmentID: '4',
          status: SERVICE_ITEM_STATUSES.APPROVED,
        },
        {
          id: '13',
          status: SERVICE_ITEM_STATUSES.REJECTED,
          mtoShipmentID: '4',
        },
        {
          id: '14',
          status: SERVICE_ITEM_STATUSES.SUBMITTED,
          mtoShipmentID: '4',
        },
      ],
    },
  ],
};

const singleReviewedPaymentRequest = {
  paymentRequests: [
    {
      id: '09474c6a-69b6-4501-8e08-670a12512e5f',
      createdAt: '2020-12-01T00:00:00.000Z',
      paymentRequestNumber: '1843-9061-1',
      status: 'REVIEWED',
      moveTaskOrderID: '1',
      moveTaskOrder: move,
      serviceItems: [
        {
          id: '09474c6a-69b6-4501-8e08-670a12512a5f',
          createdAt: '2020-12-01T00:00:00.000Z',
          mtoServiceItemID: 'f8c2f97f-99e7-4fb1-9cc4-473debd24dbc',
          priceCents: 2000001,
          status: 'APPROVED',
        },
        {
          id: '39474c6a-69b6-4501-8e08-670a12512a5f',
          createdAt: '2020-12-01T00:00:00.000Z',
          mtoServiceItemID: 'a8c2f97f-99e7-4fb1-9cc4-473debd24dbc',
          priceCents: 4000001,
          status: 'DENIED',
          rejectionReason: 'Requested amount exceeds guideline',
        },
      ],
      reviewedAt: '2020-12-01T00:00:00.000Z',
    },
  ],
  mtoShipments: [
    {
      id: '2',
      moveTaskOrderID: '1',
      status: shipmentStatuses.APPROVED,
      scheduledPickupDate: '2020-01-11T00:00:00.000Z',
      destinationAddress: { city: 'Princeton', state: 'NJ', postal_code: '08540' },
      pickupAddress: { city: 'Boston', state: 'MA', postal_code: '02101' },
      mtoServiceItems: [
        {
          id: '3',
          mtoShipmentID: '2',
          status: SERVICE_ITEM_STATUSES.APPROVED,
        },
      ],
    },
  ],
};

const emptyPaymentRequests = {
  paymentRequests: [],
  mtoShipments: [],
};

const loadingReturnValue = {
  isLoading: true,
  isError: false,
  isSuccess: false,
};

const errorReturnValue = {
  isLoading: false,
  isError: true,
  isSuccess: false,
};

function renderMovePaymentRequests(props) {
  return render(
    <MockProviders initialEntries={[`/moves/L2BKD6/payment-requests`]}>
      <MovePaymentRequests {...props} />
    </MockProviders>,
  );
}

describe('MovePaymentRequests', () => {
  describe('check loading and error component states', () => {
    it('renders the Loading Placeholder when the query is still loading', async () => {
      useMovePaymentRequestsQueries.mockReturnValue(loadingReturnValue);

      renderMovePaymentRequests(testProps);

      const h2 = await screen.getByRole('heading', { name: 'Loading, please wait...', level: 2 });
      expect(h2).toBeInTheDocument();
    });

    it('renders the Something Went Wrong component when the query errors', async () => {
      useMovePaymentRequestsQueries.mockReturnValue(errorReturnValue);

      renderMovePaymentRequests(testProps);

      const errorMessage = await screen.getByText(/Something went wrong./);
      expect(errorMessage).toBeInTheDocument();
    });
  });

  describe('with multiple payment requests', () => {
    beforeEach(() => {
      useMovePaymentRequestsQueries.mockReturnValue(multiplePaymentRequests);
    });

    it('renders without errors', () => {
      renderMovePaymentRequests(testProps);
      expect(screen.getByTestId('MovePaymentRequests')).toBeInTheDocument();
    });

    it('renders multiple payment requests', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        multiplePaymentRequests.paymentRequests.forEach((pr) => {
          expect(screen.getByText(`Payment Request ${pr.paymentRequestNumber}`)).toBeInTheDocument();
        });
      });
    });

    it('updates the pending payment request count callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setPendingPaymentRequestCount).toHaveBeenCalledWith(1);
      });
    });

    it('updates the unapproved shipments tag callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setUnapprovedShipmentCount).toHaveBeenCalledWith(1);
      });
    });

    it('updates the unapproved service items tag callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setUnapprovedServiceItemCount).toHaveBeenCalledWith(2);
      });
    });
  });

  describe('with one reviewed payment request', () => {
    beforeEach(() => {
      useMovePaymentRequestsQueries.mockReturnValue(singleReviewedPaymentRequest);
    });

    it('renders side navigation for each section', () => {
      renderMovePaymentRequests(testProps);
      const leftNav = screen.getByRole('navigation');
      expect(leftNav).toBeInTheDocument();

      const paymentRequstNavLink = within(leftNav).getByText('Payment requests');

      expect(paymentRequstNavLink.href).toContain('#payment-requests');
      expect(paymentRequstNavLink.text).toBe('Payment requests');
    });

    it('updates the pending payment request count callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setPendingPaymentRequestCount).toHaveBeenCalledWith(0);
      });
    });

    it('updates the unapproved shipment count callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setUnapprovedShipmentCount).toHaveBeenCalledWith(0);
      });
    });

    it('updates the unapproved service item count callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setUnapprovedServiceItemCount).toHaveBeenCalledWith(0);
      });
    });
  });

  describe('with no payment requests for move', () => {
    beforeEach(() => {
      useMovePaymentRequestsQueries.mockReturnValue(emptyPaymentRequests);
    });

    it('does not render side navigation for payment request section', () => {
      renderMovePaymentRequests(testProps);
      const leftNav = screen.getByRole('navigation');
      expect(leftNav).toBeInTheDocument();

      const paymentRequstNavLink = within(leftNav).queryByText('Payment requests');

      expect(paymentRequstNavLink).toBeNull();
    });

    it('renders with empty message when no payment requests exist', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(screen.getByText('No payment requests have been submitted for this move yet.')).toBeInTheDocument();
      });
    });

    it('updates the pending payment request count callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setPendingPaymentRequestCount).toHaveBeenCalledWith(0);
      });
    });

    it('updates the unapproved shipment count callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setUnapprovedShipmentCount).toHaveBeenCalledWith(0);
      });
    });

    it('updates the unapproved service item count callback', async () => {
      renderMovePaymentRequests(testProps);
      await waitFor(() => {
        expect(testProps.setUnapprovedServiceItemCount).toHaveBeenCalledWith(0);
      });
    });
  });
});
