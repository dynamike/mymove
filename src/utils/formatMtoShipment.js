import { isEmpty } from 'lodash';

import { MTOAgentType } from 'shared/constants';
import { formatSwaggerDate, parseSwaggerDate } from 'shared/formatters';

function formatAgentForDisplay(agent) {
  const agentCopy = { ...agent };
  return agentCopy;
}

function formatAgentForAPI(agent) {
  const agentCopy = { ...agent };
  Object.keys(agentCopy).forEach((key) => {
    const sanitizedKey = `${key}`;
    if (agentCopy[sanitizedKey] === '') {
      delete agentCopy[sanitizedKey];
    } else if (
      // These fields are readOnly so we don't want to send them in requests
      sanitizedKey === 'updatedAt' ||
      sanitizedKey === 'createdAt' ||
      sanitizedKey === 'mtoShipmentID'
    ) {
      delete agentCopy[sanitizedKey];
    }
  });
  return agentCopy;
}

function formatAddressForAPI(address) {
  const formattedAddress = address;

  if (formattedAddress.state) {
    formattedAddress.state = formattedAddress.state?.toUpperCase();
    delete formattedAddress.id;
    return formattedAddress;
  }

  return undefined;
}

const emptyAgentShape = {
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
};

const emptyAddressShape = {
  street_address_1: '',
  street_address_2: '',
  city: '',
  state: '',
  postal_code: '',
};

/**
 * formatMtoShipmentForDisplay converts mtoShipment data from the format API calls expect to the template format
 * @param {*} mtoShipment - (see MtoShipmentShape)
 */
export function formatMtoShipmentForDisplay({
  agents,
  shipmentType,
  requestedPickupDate,
  pickupAddress,
  requestedDeliveryDate,
  destinationAddress,
  customerRemarks,
  counselorRemarks,
  moveTaskOrderID,
  secondaryPickupAddress,
  secondaryDeliveryAddress,
}) {
  const displayValues = {
    shipmentType,
    moveTaskOrderID,
    customerRemarks: customerRemarks || '',
    counselorRemarks: counselorRemarks || '',
    pickup: {
      requestedDate: '',
      address: { ...emptyAddressShape },
      agent: { ...emptyAgentShape },
    },
    delivery: {
      requestedDate: '',
      address: { ...emptyAddressShape },
      agent: { ...emptyAgentShape },
    },
    secondaryPickup: {
      address: { ...emptyAddressShape },
    },
    secondaryDelivery: {
      address: { ...emptyAddressShape },
    },
    hasDeliveryAddress: 'no',
    hasSecondaryPickup: 'no',
    hasSecondaryDelivery: 'no',
  };

  if (agents) {
    const receivingAgent = agents.find((agent) => agent.agentType === 'RECEIVING_AGENT');
    const releasingAgent = agents.find((agent) => agent.agentType === 'RELEASING_AGENT');

    if (receivingAgent) {
      const formattedAgent = formatAgentForDisplay(receivingAgent);
      if (Object.keys(formattedAgent).length) {
        displayValues.delivery.agent = { ...emptyAgentShape, ...formattedAgent };
      }
    }
    if (releasingAgent) {
      const formattedAgent = formatAgentForDisplay(releasingAgent);
      if (Object.keys(formattedAgent).length) {
        displayValues.pickup.agent = { ...emptyAgentShape, ...formattedAgent };
      }
    }
  }

  if (pickupAddress) {
    displayValues.pickup.address = { ...emptyAddressShape, ...pickupAddress };
  }

  if (requestedPickupDate) {
    displayValues.pickup.requestedDate = parseSwaggerDate(requestedPickupDate);
  }

  if (secondaryPickupAddress) {
    displayValues.secondaryPickup.address = { ...emptyAddressShape, ...secondaryPickupAddress };
    displayValues.hasSecondaryPickup = 'yes';
  }

  if (destinationAddress) {
    displayValues.delivery.address = { ...emptyAddressShape, ...destinationAddress };
    displayValues.hasDeliveryAddress = 'yes';
  }

  if (secondaryDeliveryAddress) {
    displayValues.secondaryDelivery.address = { ...emptyAddressShape, ...secondaryDeliveryAddress };
    displayValues.hasSecondaryDelivery = 'yes';
  }

  if (requestedDeliveryDate) {
    displayValues.delivery.requestedDate = parseSwaggerDate(requestedDeliveryDate);
  }

  return displayValues;
}

/**
 * formatMtoShipmentForAPI converts mtoShipment data from the template format to the format API calls expect
 * @param {*} param - unnamed object representing various mtoShipment data parts
 */
export function formatMtoShipmentForAPI({
  moveId,
  shipmentType,
  pickup,
  delivery,
  customerRemarks,
  counselorRemarks,
  secondaryPickup,
  secondaryDelivery,
}) {
  const formattedMtoShipment = {
    moveTaskOrderID: moveId,
    shipmentType,
    customerRemarks,
    counselorRemarks,
    agents: [],
  };

  if (pickup?.requestedDate && pickup.requestedDate !== '') {
    formattedMtoShipment.requestedPickupDate = formatSwaggerDate(pickup.requestedDate);
    formattedMtoShipment.pickupAddress = formatAddressForAPI(pickup.address);

    if (pickup.agent) {
      const formattedAgent = formatAgentForAPI(pickup.agent);
      if (!isEmpty(formattedAgent)) {
        formattedMtoShipment.agents.push({ ...formattedAgent, agentType: MTOAgentType.RELEASING });
      }
    }
  }

  if (delivery?.requestedDate && delivery.requestedDate !== '') {
    formattedMtoShipment.requestedDeliveryDate = formatSwaggerDate(delivery.requestedDate);

    if (delivery.address) {
      formattedMtoShipment.destinationAddress = formatAddressForAPI(delivery.address);
    }

    if (delivery.agent) {
      const formattedAgent = formatAgentForAPI(delivery.agent);
      if (!isEmpty(formattedAgent)) {
        formattedMtoShipment.agents.push({ ...formattedAgent, agentType: MTOAgentType.RECEIVING });
      }
    }
  }

  if (secondaryPickup?.address) {
    formattedMtoShipment.secondaryPickupAddress = formatAddressForAPI(secondaryPickup.address);
  }

  if (secondaryDelivery?.address) {
    formattedMtoShipment.secondaryDeliveryAddress = formatAddressForAPI(secondaryDelivery.address);
  }

  if (!formattedMtoShipment.agents?.length) {
    formattedMtoShipment.agents = undefined;
  }

  return formattedMtoShipment;
}

export default { formatMtoShipmentForAPI, formatMtoShipmentForDisplay };
