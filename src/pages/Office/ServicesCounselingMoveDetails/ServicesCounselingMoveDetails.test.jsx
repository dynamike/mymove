/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import { generatePath } from 'react-router';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

import ServicesCounselingMoveDetails from './ServicesCounselingMoveDetails';

import MOVE_STATUSES from 'constants/moves';
import { ORDERS_TYPE, ORDERS_TYPE_DETAILS } from 'constants/orders';
import { servicesCounselingRoutes } from 'constants/routes';
import { useMoveDetailsQueries } from 'hooks/queries';
import { formatDate } from 'shared/dates';
import { MockProviders, renderWithRouter } from 'testUtils';

const mockRequestedMoveCode = 'LR4T8V';

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useParams: jest.fn().mockReturnValue({ moveCode: 'LR4T8V' }),
}));

jest.mock('hooks/queries', () => ({
  useMoveDetailsQueries: jest.fn(),
}));

const mtoShipments = [
  {
    customerRemarks: 'please treat gently',
    counselorRemarks: 'all good',
    destinationAddress: {
      city: 'Fairfield',
      country: 'US',
      id: '672ff379-f6e3-48b4-a87d-796713f8f997',
      postal_code: '94535',
      state: 'CA',
      street_address_1: '987 Any Avenue',
      street_address_2: 'P.O. Box 9876',
      street_address_3: 'c/o Some Person',
    },
    eTag: 'MjAyMC0wNi0xMFQxNTo1ODowMi40MDQwMzFa',
    id: 'ce01a5b8-9b44-4511-8a8d-edb60f2a4aee',
    moveTaskOrderID: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
    pickupAddress: {
      city: 'Beverly Hills',
      country: 'US',
      eTag: 'MjAyMC0wNi0xMFQxNTo1ODowMi4zODQ3Njla',
      id: '1686751b-ab36-43cf-b3c9-c0f467d13c19',
      postal_code: '90210',
      state: 'CA',
      street_address_1: '123 Any Street',
      street_address_2: 'P.O. Box 12345',
      street_address_3: 'c/o Some Person',
    },
    secondaryPickupAddress: {
      city: 'Los Angeles',
      country: 'US',
      eTag: 'MjAyMC0wNi0xMFQxNTo1ODowMi4zODQ3Njla',
      id: 'b941a74a-e77e-4575-bea3-e7e01b226422',
      postal_code: '90222',
      state: 'CA',
      street_address_1: '456 Any Street',
      street_address_2: 'P.O. Box 67890',
      street_address_3: 'c/o A Friendly Person',
    },
    secondaryDeliveryAddress: {
      city: 'Beverly Hills',
      country: 'US',
      eTag: 'MjAyMC0wNi0xMFQxNTo1ODowMi4zODQ3Njla',
      id: '1686751b-ab36-43cf-eeee-c0f467d13c19',
      postal_code: '90215',
      state: 'CA',
      street_address_1: '123 Any Street',
      street_address_2: 'P.O. Box 12345',
      street_address_3: 'c/o Some Person',
    },
    requestedPickupDate: '2020-06-04',
    scheduledPickupDate: '2020-06-05',
    shipmentType: 'HHG',
    status: 'SUBMITTED',
    updatedAt: '2020-05-10T15:58:02.404031Z',
  },
  {
    customerRemarks: 'do not drop!',
    counselorRemarks: 'looks good',
    destinationAddress: {
      city: 'Fairfield',
      country: 'US',
      id: '672ff379-f6e3-48b4-a87d-752463f8f997',
      postal_code: '94534',
      state: 'CA',
      street_address_1: '111 Everywhere',
      street_address_2: 'Apt #1',
      street_address_3: '',
    },
    eTag: 'MjAyMC0wNi0xMFQxNTo1ODowMi40MDQwMzFa',
    id: 'ce01a5b8-9b44-8799-8a8d-edb60f2a4aee',
    moveTaskOrderID: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
    pickupAddress: {
      city: 'Austin',
      country: 'US',
      eTag: 'MjAyMC0wNi0xMFQxNTo1ODowMi4zODQ3Njla',
      id: '1686751b-ab36-43cf-b3c9-c0f467d13c55',
      postal_code: '78712',
      state: 'TX',
      street_address_1: '888 Lucky Street',
      street_address_2: '#4',
      street_address_3: 'c/o rabbit',
    },
    requestedPickupDate: '2020-06-05',
    scheduledPickupDate: '2020-06-06',
    shipmentType: 'HHG',
    status: 'SUBMITTED',
    updatedAt: '2020-05-15T15:58:02.404031Z',
  },
];

const newMoveDetailsQuery = {
  move: {
    id: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
    ordersId: '1',
    status: MOVE_STATUSES.NEEDS_SERVICE_COUNSELING,
  },
  order: {
    id: '1',
    originDutyStation: {
      address: {
        street_address_1: '',
        city: 'Fort Knox',
        state: 'KY',
        postal_code: '40121',
      },
    },
    destinationDutyStation: {
      address: {
        street_address_1: '',
        city: 'Fort Irwin',
        state: 'CA',
        postal_code: '92310',
      },
    },
    customer: {
      agency: 'ARMY',
      backup_contact: {
        email: 'email@example.com',
        name: 'name',
        phone: '555-555-5555',
      },
      current_address: {
        city: 'Beverly Hills',
        country: 'US',
        eTag: 'MjAyMS0wMS0yMVQxNTo0MTozNS41Mzg0Njha',
        id: '3a5f7cf2-6193-4eb3-a244-14d21ca05d7b',
        postal_code: '90210',
        state: 'CA',
        street_address_1: '123 Any Street',
        street_address_2: 'P.O. Box 12345',
        street_address_3: 'c/o Some Person',
      },
      dodID: '6833908165',
      eTag: 'MjAyMS0wMS0yMVQxNTo0MTozNS41NjAzNTJa',
      email: 'combo@ppm.hhg',
      first_name: 'Submitted',
      id: 'f6bd793f-7042-4523-aa30-34946e7339c9',
      last_name: 'Ppmhhg',
      phone: '555-555-5555',
    },
    entitlement: {
      authorizedWeight: 8000,
      dependentsAuthorized: true,
      eTag: 'MjAyMS0wMS0yMVQxNTo0MTozNS41NzgwMzda',
      id: 'e0fefe58-0710-40db-917b-5b96567bc2a8',
      nonTemporaryStorage: true,
      privatelyOwnedVehicle: true,
      proGearWeight: 2000,
      proGearWeightSpouse: 500,
      storageInTransit: 2,
      totalDependents: 1,
      totalWeight: 8000,
    },
    order_number: 'ORDER3',
    order_type: ORDERS_TYPE.PERMANENT_CHANGE_OF_STATION,
    order_type_detail: ORDERS_TYPE_DETAILS.HHG_PERMITTED,
    tac: '9999',
  },
  mtoShipments,
  mtoServiceItems: [],
  mtoAgents: [],
  isLoading: false,
  isError: false,
  isSuccess: true,
};

const counselingCompletedMoveDetailsQuery = {
  ...newMoveDetailsQuery,
  move: {
    id: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
    ordersId: '1',
    status: MOVE_STATUSES.SERVICE_COUNSELING_COMPLETED,
  },
};

const detailsURL = generatePath(servicesCounselingRoutes.MOVE_VIEW_PATH, { moveCode: mockRequestedMoveCode });

const renderMockedComponent = (props) => {
  return render(
    <MockProviders initialEntries={[detailsURL]}>
      <ServicesCounselingMoveDetails {...props} />
    </MockProviders>,
  );
};

const mockedComponent = (
  <MockProviders initialEntries={[detailsURL]}>
    <ServicesCounselingMoveDetails />
  </MockProviders>
);

const loadingReturnValue = {
  ...newMoveDetailsQuery,
  isLoading: true,
  isError: false,
  isSuccess: false,
};

const errorReturnValue = {
  ...newMoveDetailsQuery,
  isLoading: false,
  isError: true,
  isSuccess: false,
};

describe('MoveDetails page', () => {
  describe('check loading and error component states', () => {
    it('renders the Loading Placeholder when the query is still loading', async () => {
      useMoveDetailsQueries.mockReturnValue(loadingReturnValue);

      render(mockedComponent);

      const h2 = await screen.getByRole('heading', { name: 'Loading, please wait...', level: 2 });
      expect(h2).toBeInTheDocument();
    });

    it('renders the Something Went Wrong component when the query errors', async () => {
      useMoveDetailsQueries.mockReturnValue(errorReturnValue);

      render(mockedComponent);

      const errorMessage = await screen.getByText(/Something went wrong./);
      expect(errorMessage).toBeInTheDocument();
    });
  });

  describe('Basic rendering', () => {
    it('renders the h1', async () => {
      useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

      render(mockedComponent);

      expect(await screen.findByRole('heading', { name: 'Move details', level: 1 })).toBeInTheDocument();
    });

    /* eslint-disable camelcase */
    it('renders shipments info', async () => {
      useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

      render(mockedComponent);

      expect(await screen.findByRole('heading', { name: 'Shipments', level: 2 })).toBeInTheDocument();

      expect(screen.getAllByRole('heading', { name: 'HHG', level: 3 }).length).toBe(2);

      const moveDateTerms = screen.getAllByText('Requested move date');

      expect(moveDateTerms.length).toBe(2);

      for (let i = 0; i < moveDateTerms.length; i += 1) {
        expect(moveDateTerms[i].nextElementSibling.textContent).toBe(
          formatDate(newMoveDetailsQuery.mtoShipments[i].requestedPickupDate, 'DD MMM YYYY'),
        );
      }

      const originAddressTerms = screen.getAllByText('Origin address');

      expect(originAddressTerms.length).toBe(2);

      for (let i = 0; i < 2; i += 1) {
        const { street_address_1, city, state, postal_code } = newMoveDetailsQuery.mtoShipments[i].pickupAddress;

        const addressText = originAddressTerms[i].nextElementSibling.textContent;

        expect(addressText).toContain(street_address_1);
        expect(addressText).toContain(city);
        expect(addressText).toContain(state);
        expect(addressText).toContain(postal_code);
      }

      const secondAddressTerms = screen.getAllByText('Second pickup address');

      expect(secondAddressTerms.length).toBe(1);

      for (let i = 0; i < 1; i += 1) {
        const { street_address_1, city, state, postal_code } =
          newMoveDetailsQuery.mtoShipments[i].secondaryPickupAddress;

        const addressText = secondAddressTerms[0].nextElementSibling.textContent;

        expect(addressText).toContain(street_address_1);
        expect(addressText).toContain(city);
        expect(addressText).toContain(state);
        expect(addressText).toContain(postal_code);
      }

      const destinationAddressTerms = screen.getAllByText('Destination address');

      expect(destinationAddressTerms.length).toBe(2);

      for (let i = 0; i < destinationAddressTerms.length; i += 1) {
        const { street_address_1, city, state, postal_code } = newMoveDetailsQuery.mtoShipments[i].destinationAddress;

        const addressText = destinationAddressTerms[i].nextElementSibling.textContent;

        expect(addressText).toContain(street_address_1);
        expect(addressText).toContain(city);
        expect(addressText).toContain(state);
        expect(addressText).toContain(postal_code);
      }

      const secondDestinationAddressTerms = screen.getAllByText('Second destination address');

      // This is not a required field, and only one of our shipments has it filled out:
      expect(secondDestinationAddressTerms.length).toBe(1);

      const { street_address_1, city, state, postal_code } =
        newMoveDetailsQuery.mtoShipments[0].secondaryDeliveryAddress;
      const addressText = secondDestinationAddressTerms[0].nextElementSibling.textContent;

      expect(addressText).toContain(street_address_1);
      expect(addressText).toContain(city);
      expect(addressText).toContain(state);
      expect(addressText).toContain(postal_code);

      const counselorRemarksTerms = screen.getAllByText('Counselor remarks');

      expect(counselorRemarksTerms.length).toBe(2);

      for (let i = 0; i < counselorRemarksTerms.length; i += 1) {
        expect(counselorRemarksTerms[i].nextElementSibling.textContent).toBe(
          newMoveDetailsQuery.mtoShipments[i].counselorRemarks || '—',
        );
      }
    });

    it('renders shipments info even if destination address is missing', async () => {
      const moveDetailsQuery = {
        ...newMoveDetailsQuery,
        mtoShipments: [
          // Want to create a "new" mtoShipment to be able to delete things without messing up existing tests
          { ...newMoveDetailsQuery.mtoShipments[0] },
          newMoveDetailsQuery.mtoShipments[1],
        ],
      };

      delete moveDetailsQuery.mtoShipments[0].destinationAddress;

      useMoveDetailsQueries.mockReturnValue(moveDetailsQuery);

      render(mockedComponent);

      const destinationAddressTerms = screen.getAllByText('Destination address');

      expect(destinationAddressTerms.length).toBe(2);

      expect(destinationAddressTerms[0].nextElementSibling.textContent).toBe(
        moveDetailsQuery.order.destinationDutyStation.address.postal_code,
      );

      const { street_address_1, city, state, postal_code } = moveDetailsQuery.mtoShipments[1].destinationAddress;

      const addressText = destinationAddressTerms[1].nextElementSibling.textContent;

      expect(addressText).toContain(street_address_1);
      expect(addressText).toContain(city);
      expect(addressText).toContain(state);
      expect(addressText).toContain(postal_code);
    });
    /* eslint-enable camelcase */

    it('renders customer info', async () => {
      useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

      render(mockedComponent);

      expect(await screen.findByRole('heading', { name: 'Customer info', level: 2 })).toBeInTheDocument();
    });

    it('renders customer edit alert', () => {
      renderMockedComponent({ customerEditAlert: { alertType: 'success', message: 'great success!' } });
      expect(screen.getByText('great success!')).toBeInTheDocument();
    });

    describe('new move - needs service counseling', () => {
      it('submit move details button is on page', async () => {
        useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

        render(mockedComponent);

        expect(await screen.findByRole('button', { name: 'Submit move details' })).toBeInTheDocument();
      });

      it('submit move details button is disabled when there are no shipments', async () => {
        useMoveDetailsQueries.mockReturnValue({ ...newMoveDetailsQuery, mtoShipments: [] });

        render(mockedComponent);

        expect(await screen.findByRole('button', { name: 'Submit move details' })).toBeInTheDocument();
        expect(await screen.findByRole('button', { name: 'Submit move details' })).toBeDisabled();
      });

      it('submit move details button is disabled when all shipments are deleted', async () => {
        const deletedMtoShipments = mtoShipments.map((shipment) => ({ ...shipment, deletedAt: new Date() }));
        useMoveDetailsQueries.mockReturnValue({
          ...newMoveDetailsQuery,
          mtoShipments: deletedMtoShipments,
        });

        render(mockedComponent);

        expect(await screen.findByRole('button', { name: 'Submit move details' })).toBeInTheDocument();
        expect(await screen.findByRole('button', { name: 'Submit move details' })).toBeDisabled();
      });

      it('submit move details button is not disabled when some shipments are deleted', async () => {
        const deletedMtoShipments = mtoShipments.map((shipment, index) => {
          if (index > 0) {
            return { ...shipment, deletedAt: new Date() };
          }
          return shipment;
        });
        useMoveDetailsQueries.mockReturnValue({
          ...newMoveDetailsQuery,
          mtoShipments: deletedMtoShipments,
        });

        render(mockedComponent);

        expect(await screen.findByRole('button', { name: 'Submit move details' })).toBeInTheDocument();
        expect(await screen.findByRole('button', { name: 'Submit move details' })).not.toBeDisabled();
      });

      it('renders the Orders Definition List', async () => {
        useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

        render(mockedComponent);

        expect(await screen.findByRole('heading', { name: 'Orders', level: 2 })).toBeInTheDocument();
        expect(screen.getByText('Current duty station')).toBeInTheDocument();
      });

      it('renders the Allowances Table', async () => {
        useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

        render(mockedComponent);

        expect(await screen.findByRole('heading', { name: 'Allowances', level: 2 })).toBeInTheDocument();
        expect(screen.getByText('Branch, rank')).toBeInTheDocument();
      });

      it('allows the service counselor to use the modal as expected', async () => {
        useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

        render(mockedComponent);

        const submitButton = await screen.findByRole('button', { name: 'Submit move details' });

        userEvent.click(submitButton);

        expect(await screen.findByRole('heading', { name: 'Are you sure?', level: 2 }));

        const modalSubmitButton = screen.getByRole('button', { name: 'Yes, submit' });

        userEvent.click(modalSubmitButton);

        expect(screen.queryByRole('heading', { name: 'Are you sure?', level: 2 }));
      });

      it.each([
        ['Add a new shipment', servicesCounselingRoutes.SHIPMENT_ADD_PATH],
        ['View and edit orders', servicesCounselingRoutes.ORDERS_EDIT_PATH],
        ['Edit allowances', servicesCounselingRoutes.ALLOWANCES_EDIT_PATH],
        ['Edit customer info', servicesCounselingRoutes.CUSTOMER_INFO_EDIT_PATH],
      ])('shows the "%s" link as expected: %s', async (linkText, route) => {
        useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

        const { history } = renderWithRouter(<ServicesCounselingMoveDetails />, { route: detailsURL });

        const link = await screen.findByRole('link', { name: linkText });

        expect(link).toBeInTheDocument();

        userEvent.click(link);

        const path = generatePath(route, {
          moveCode: mockRequestedMoveCode,
        });

        await waitFor(() => {
          expect(history.location.pathname).toEqual(path);
        });
      });

      it('shows the edit shipment buttons', async () => {
        useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

        render(mockedComponent);

        const shipmentEditButtons = await screen.findAllByRole('button', { name: 'Edit shipment' });

        expect(shipmentEditButtons.length).toBe(2);

        for (let i = 0; i < shipmentEditButtons.length; i += 1) {
          expect(shipmentEditButtons[i].getAttribute('data-testid')).toBe(
            generatePath(servicesCounselingRoutes.SHIPMENT_EDIT_PATH, {
              moveCode: mockRequestedMoveCode,
              shipmentId: newMoveDetailsQuery.mtoShipments[i].id,
            }),
          );
        }
      });

      it('shows the customer and counselor remarks', async () => {
        useMoveDetailsQueries.mockReturnValue(newMoveDetailsQuery);

        render(mockedComponent);

        const customerRemarks1 = await screen.findByText('please treat gently');
        const customerRemarks2 = await screen.findByText('do not drop!');

        const counselorRemarks1 = await screen.findByText('all good');
        const counselorRemarks2 = await screen.findByText('looks good');

        expect(customerRemarks1).toBeInTheDocument();
        expect(customerRemarks2).toBeInTheDocument();
        expect(counselorRemarks1).toBeInTheDocument();
        expect(counselorRemarks2).toBeInTheDocument();
      });
    });

    describe('service counseling completed', () => {
      it('hides submit and view/edit buttons/links', async () => {
        useMoveDetailsQueries.mockReturnValue(counselingCompletedMoveDetailsQuery);

        render(mockedComponent);

        expect(screen.queryByRole('button', { name: 'Submit move details' })).not.toBeInTheDocument();
        expect(screen.queryByRole('link', { name: 'Add a new shipment' })).not.toBeInTheDocument();
        expect(screen.queryByRole('button', { name: 'Edit shipment' })).not.toBeInTheDocument();
        expect(screen.queryByRole('link', { name: 'View and edit orders' })).not.toBeInTheDocument();
        expect(screen.queryByRole('link', { name: 'Edit allowances' })).not.toBeInTheDocument();
        expect(screen.queryByRole('link', { name: 'Edit customer info' })).not.toBeInTheDocument();
      });
    });
  });
});
