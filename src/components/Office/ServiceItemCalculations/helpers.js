import { SERVICE_ITEM_CALCULATION_LABELS, SERVICE_ITEM_CODES, SERVICE_ITEM_PARAM_KEYS } from 'constants/serviceItems';
import { LONGHAUL_MIN_DISTANCE } from 'constants/shipments';
import { formatWeight, formatCents, toDollarString } from 'shared/formatters';
import { formatDate } from 'shared/dates';
import { formatWeightCWTFromLbs, formatDollarFromMillicents } from 'utils/formatters';

const calculation = (value, label, ...details) => {
  return {
    value,
    label,
    details: [...details],
  };
};

const getParamValue = (key, params) => {
  return params?.find((param) => param?.key === key)?.value;
};

const peak = (params) => {
  return `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.IsPeak]} ${
    getParamValue(SERVICE_ITEM_PARAM_KEYS.IsPeak, params)?.toLowerCase() === 'true' ? 'peak' : 'non-peak'
  }`;
};

const serviceAreaOrigin = (params) => {
  return `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ServiceAreaOrigin]}: ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ServiceAreaOrigin,
    params,
  )}`;
};

const serviceAreaDest = (params) => {
  return `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ServiceAreaDest]}: ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ServiceAreaDest,
    params,
  )}`;
};

const requestedPickupDate = (params) => {
  return `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.RequestedPickupDate]}: ${formatDate(
    getParamValue(SERVICE_ITEM_PARAM_KEYS.RequestedPickupDate, params),
    'DD MMM YYYY',
  )}`;
};

const cratingDate = (params) => {
  return `${SERVICE_ITEM_CALCULATION_LABELS.CratingDate}: ${formatDate(
    getParamValue(SERVICE_ITEM_PARAM_KEYS.RequestedPickupDate, params),
    'DD MMM YYYY',
  )}`;
};

const unCratingDate = (params) => {
  return `${SERVICE_ITEM_CALCULATION_LABELS.UncratingDate}: ${formatDate(
    getParamValue(SERVICE_ITEM_PARAM_KEYS.RequestedPickupDate, params),
    'DD MMM YYYY',
  )}`;
};

const getPriceRateOrFactor = (params) => {
  return getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params) || '';
};

// billable weight calculation
const billableWeight = (params) => {
  const value = formatWeightCWTFromLbs(getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightBilledActual, params));
  const label = SERVICE_ITEM_CALCULATION_LABELS.BillableWeight;

  const weightBilledActualDetail = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.WeightBilledActual]
  }: ${formatWeight(parseInt(getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightBilledActual, params), 10))}`;

  const weightEstimated = getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightEstimated, params);
  const weightEstimatedDetail = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.WeightEstimated]}: ${
    weightEstimated ? formatWeight(parseInt(getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightEstimated, params), 10)) : ''
  }`;
  return calculation(value, label, weightBilledActualDetail, weightEstimatedDetail);
};

const shuttleBillableWeight = (params) => {
  const value = formatWeightCWTFromLbs(getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightBilledActual, params));
  const label = SERVICE_ITEM_CALCULATION_LABELS.BillableWeight;

  const weightBilledActualDetail = `${SERVICE_ITEM_CALCULATION_LABELS.ShuttleWeight}: ${formatWeight(
    parseInt(getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightBilledActual, params), 10),
  )}`;

  const weightEstimated = getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightEstimated, params);
  const weightEstimatedDetail = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.WeightEstimated]}: ${
    weightEstimated ? formatWeight(parseInt(getParamValue(SERVICE_ITEM_PARAM_KEYS.WeightEstimated, params), 10)) : ''
  }`;
  return calculation(value, label, weightBilledActualDetail, weightEstimatedDetail);
};

// display the first 3 digits of the ZIP code
const mileageFirstThreeZip = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.DistanceZip3, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.Mileage;
  const detail = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipPickupAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipPickupAddress,
    params,
  )?.slice(0, 3)} to ${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipDestAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipDestAddress,
    params,
  )?.slice(0, 3)}`;

  return calculation(value, label, detail);
};

const mileageZip5 = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.DistanceZip5, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.Mileage;
  const detail = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipPickupAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipPickupAddress,
    params,
  )} to ${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipDestAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipDestAddress,
    params,
  )}`;

  return calculation(value, label, detail);
};

const mileageZipSITOrigin = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.DistanceZipSITOrigin, params);

  const label = SERVICE_ITEM_CALCULATION_LABELS.Mileage;
  const detail = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipPickupAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipSITOriginHHGOriginalAddress,
    params,
  )} to ${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipDestAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipSITOriginHHGActualAddress,
    params,
  )}`;
  return calculation(value, label, detail);
};

const baselineLinehaulPrice = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.BaselineLinehaulPrice;

  return calculation(value, label, peak(params), serviceAreaOrigin(params), requestedPickupDate(params));
};

const baselineShorthaulPrice = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.BaselineShorthaulPrice;

  return calculation(value, label, peak(params), serviceAreaOrigin(params), requestedPickupDate(params));
};

const dddSITmileageZip5 = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.DistanceZipSITDest, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.Mileage;
  const detail = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipDestAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipDestAddress,
    params,
  )} to ${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ZipSITDestHHGFinalAddress]} ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ZipSITDestHHGFinalAddress,
    params,
  )}`;

  return calculation(value, label, detail);
};

// There is no param representing the orgin price as available in the re_domestic_service_area_prices table
// A param to return the service schedule is also not being created
const originPrice = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.OriginPrice;

  return calculation(value, label, serviceAreaOrigin(params), requestedPickupDate(params), peak(params));
};

const shuttleOriginPriceDomestic = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.OriginPrice;

  const serviceSchedule = `${SERVICE_ITEM_CALCULATION_LABELS.ServiceSchedule}: ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ServicesScheduleOrigin,
    params,
  )}`;

  const pickupDate = `${SERVICE_ITEM_CALCULATION_LABELS.PickupDate}: ${formatDate(
    getParamValue(SERVICE_ITEM_PARAM_KEYS.RequestedPickupDate, params),
    'DD MMM YYYY',
  )}`;

  return calculation(value, label, serviceSchedule, pickupDate, SERVICE_ITEM_CALCULATION_LABELS.Domestic);
};

// There is no param representing the destination price as available in the re_domestic_service_area_prices table
// A param to return the service schedule is also not being created
const destinationPrice = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.DestinationPrice;

  return calculation(value, label, serviceAreaDest(params), requestedPickupDate(params), peak(params));
};

const shuttleDestinationPriceDomestic = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.DestinationPrice;

  const serviceSchedule = `${SERVICE_ITEM_CALCULATION_LABELS.ServiceSchedule}: ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ServicesScheduleDest,
    params,
  )}`;

  const deliveryDate = `${SERVICE_ITEM_CALCULATION_LABELS.DeliveryDate}: ${formatDate(
    getParamValue(SERVICE_ITEM_PARAM_KEYS.RequestedPickupDate, params),
    'DD MMM YYYY',
  )}`;

  return calculation(value, label, serviceSchedule, deliveryDate, SERVICE_ITEM_CALCULATION_LABELS.Domestic);
};

const priceEscalationFactor = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.EscalationCompounded, params)
    ? getParamValue(SERVICE_ITEM_PARAM_KEYS.EscalationCompounded, params)
    : '';
  const label = SERVICE_ITEM_CALCULATION_LABELS.PriceEscalationFactor;

  const contractYearName = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ContractYearName]}: ${
    getParamValue(SERVICE_ITEM_PARAM_KEYS.ContractYearName, params) || ''
  }`;

  return calculation(value, label, contractYearName);
};

const priceEscalationFactorWithoutContractYear = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.EscalationCompounded, params)
    ? getParamValue(SERVICE_ITEM_PARAM_KEYS.EscalationCompounded, params)
    : '';
  const label = SERVICE_ITEM_CALCULATION_LABELS.PriceEscalationFactor;

  return calculation(value, label);
};

const fuelSurchargePrice = (params) => {
  // to get the Fuel surcharge price (per mi), multiply FSCWeightBasedDistanceMultiplier by DistanceZip3
  // which gets the dollar value
  const value = parseFloat(
    String(
      getParamValue(SERVICE_ITEM_PARAM_KEYS.FSCWeightBasedDistanceMultiplier, params) *
        getParamValue(SERVICE_ITEM_PARAM_KEYS.DistanceZip3, params),
    ),
  ).toFixed(2);
  const label = SERVICE_ITEM_CALCULATION_LABELS.FuelSurchargePrice;

  const eiaFuelPrice = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.EIAFuelPrice]
  }: ${formatDollarFromMillicents(getParamValue(SERVICE_ITEM_PARAM_KEYS.EIAFuelPrice, params))}`;

  const fscWeightBasedDistanceMultiplier = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.FSCWeightBasedDistanceMultiplier]
  }: ${getParamValue(SERVICE_ITEM_PARAM_KEYS.FSCWeightBasedDistanceMultiplier, params)}`;

  const actualPickupDate = `${SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ActualPickupDate]}: ${formatDate(
    getParamValue(SERVICE_ITEM_PARAM_KEYS.ActualPickupDate, params),
    'DD MMM YYYY',
  )}`;

  return calculation(value, label, eiaFuelPrice, fscWeightBasedDistanceMultiplier, actualPickupDate);
};

const packPrice = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.PackPrice;
  const originServiceSchedule = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ServicesScheduleOrigin]
  }: ${getParamValue(SERVICE_ITEM_PARAM_KEYS.ServicesScheduleOrigin, params)}`;

  return calculation(value, label, originServiceSchedule, requestedPickupDate(params), peak(params));
};

const unpackPrice = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.UnpackPrice;
  const destServiceSchedule = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.ServicesScheduleDest]
  }: ${getParamValue(SERVICE_ITEM_PARAM_KEYS.ServicesScheduleDest, params)}`;

  return calculation(value, label, destServiceSchedule, requestedPickupDate(params), peak(params));
};

const additionalDayOriginSITPrice = (params) => {
  const value = getPriceRateOrFactor(params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.AdditionalDaySITPrice;

  return calculation(value, label, serviceAreaOrigin(params), requestedPickupDate(params), peak(params));
};

const additionalDayDestinationSITPrice = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.AdditionalDaySITPrice;

  return calculation(value, label, serviceAreaDest(params), requestedPickupDate(params), peak(params));
};

const sitDeliveryPrice = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.SITDeliveryPrice;

  const sitScheduleDestination = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.SITScheduleDest]
  }: ${getParamValue(SERVICE_ITEM_PARAM_KEYS.SITScheduleDest, params)}`;

  return calculation(value, label, sitScheduleDestination, requestedPickupDate(params), peak(params));
};

const sitDeliveryPriceShorthaulDifferentZIP3 = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.SITDeliveryPrice;

  const sitScheduleDestination = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.SITScheduleDest]
  }: ${getParamValue(SERVICE_ITEM_PARAM_KEYS.SITScheduleDest, params)}`;

  return calculation(value, label, sitScheduleDestination, requestedPickupDate(params), peak(params), '<=50 miles');
};

const daysInSIT = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.NumberDaysSIT, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.DaysInSIT;

  return calculation(value, label);
};

const pickupSITPrice = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.PickupSITPrice;

  const originSITSchedule = `${
    SERVICE_ITEM_CALCULATION_LABELS[SERVICE_ITEM_PARAM_KEYS.SITScheduleOrigin]
  }: ${getParamValue(SERVICE_ITEM_PARAM_KEYS.SITScheduleOrigin, params)}`;

  return calculation(value, label, originSITSchedule, requestedPickupDate(params), peak(params));
};

const cratingPrice = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.CratingPrice;

  const serviceSchedule = `${SERVICE_ITEM_CALCULATION_LABELS.ServiceSchedule}: ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ServicesScheduleOrigin,
    params,
  )}`;

  return calculation(value, label, serviceSchedule, cratingDate(params), SERVICE_ITEM_CALCULATION_LABELS.Domestic);
};

const unCratingPrice = (params) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.PriceRateOrFactor, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.UncratingPrice;

  const serviceSchedule = `${SERVICE_ITEM_CALCULATION_LABELS.ServiceSchedule}: ${getParamValue(
    SERVICE_ITEM_PARAM_KEYS.ServicesScheduleDest,
    params,
  )}`;

  return calculation(value, label, serviceSchedule, unCratingDate(params), SERVICE_ITEM_CALCULATION_LABELS.Domestic);
};

const cratingSize = (params, mtoParams) => {
  const value = getParamValue(SERVICE_ITEM_PARAM_KEYS.CubicFeetBilled, params);
  const length = getParamValue(SERVICE_ITEM_PARAM_KEYS.DimensionLength, params);
  const height = getParamValue(SERVICE_ITEM_PARAM_KEYS.DimensionHeight, params);
  const width = getParamValue(SERVICE_ITEM_PARAM_KEYS.DimensionWidth, params);
  const label = SERVICE_ITEM_CALCULATION_LABELS.CubicFeetBilled;

  const description = `${SERVICE_ITEM_CALCULATION_LABELS.Description}: ${mtoParams.description}`;

  const formattedDimensions = `${SERVICE_ITEM_CALCULATION_LABELS.Dimensions}: ${length}x${width}x${height} in`;

  return calculation(value, label, description, formattedDimensions);
};

// totalAmountRequested is not a service item param
const totalAmountRequested = (totalAmount) => {
  const value = toDollarString(formatCents(totalAmount));
  const label = SERVICE_ITEM_CALCULATION_LABELS.TotalAmountRequested;
  const detail = '';

  return calculation(value, label, detail);
};

const makeCalculations = (itemCode, totalAmount, params, mtoParams) => {
  let result = [];

  switch (itemCode) {
    case SERVICE_ITEM_CODES.DDDSIT: {
      const mileage = getParamValue(SERVICE_ITEM_PARAM_KEYS.DistanceZipSITDest, params);
      const startZip = getParamValue(SERVICE_ITEM_PARAM_KEYS.ZipDestAddress, params)?.slice(0, 3);
      const endZip = getParamValue(SERVICE_ITEM_PARAM_KEYS.ZipSITDestHHGFinalAddress, params)?.slice(0, 3);
      // Mileage does not factor into the pricing for distances less than 50 miles and non-matching
      // zip3, so we won't display mileage
      if (mileage <= LONGHAUL_MIN_DISTANCE && startZip !== endZip) {
        result = [
          billableWeight(params),
          sitDeliveryPriceShorthaulDifferentZIP3(params), // Display under mileage threshold
          priceEscalationFactor(params),
          totalAmountRequested(totalAmount),
        ];
      } else {
        result = [
          billableWeight(params),
          dddSITmileageZip5(params),
          sitDeliveryPrice(params),
          priceEscalationFactor(params),
          totalAmountRequested(totalAmount),
        ];
      }
      break;
    }
    // Domestic longhaul
    case SERVICE_ITEM_CODES.DLH:
      result = [
        billableWeight(params),
        mileageFirstThreeZip(params),
        baselineLinehaulPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Fuel surcharge
    case SERVICE_ITEM_CODES.FSC:
      result = [
        billableWeight(params),
        mileageFirstThreeZip(params),
        fuelSurchargePrice(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic origin price
    case SERVICE_ITEM_CODES.DOP:
      result = [
        billableWeight(params),
        originPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic origin 1st day SIT
    case SERVICE_ITEM_CODES.DOFSIT:
      result = [
        billableWeight(params),
        originPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic destination 1st day SIT
    case SERVICE_ITEM_CODES.DDFSIT:
      result = [
        billableWeight(params),
        destinationPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic packing
    case SERVICE_ITEM_CODES.DPK:
      result = [
        billableWeight(params),
        packPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic shorthaul
    case SERVICE_ITEM_CODES.DSH:
      result = [
        billableWeight(params),
        mileageZip5(params),
        baselineShorthaulPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic destination
    case SERVICE_ITEM_CODES.DDP:
      result = [
        billableWeight(params),
        destinationPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic origin additional SIT
    case SERVICE_ITEM_CODES.DOASIT:
      result = [
        billableWeight(params),
        daysInSIT(params),
        additionalDayOriginSITPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    case SERVICE_ITEM_CODES.DOPSIT:
      result = [
        billableWeight(params),
        mileageZipSITOrigin(params),
        pickupSITPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic origin shuttle service
    case SERVICE_ITEM_CODES.DOSHUT:
      result = [
        shuttleBillableWeight(params),
        shuttleOriginPriceDomestic(params),
        priceEscalationFactorWithoutContractYear(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    case SERVICE_ITEM_CODES.DDASIT:
      result = [
        billableWeight(params),
        daysInSIT(params),
        additionalDayDestinationSITPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic unpacking
    case SERVICE_ITEM_CODES.DUPK:
      result = [
        billableWeight(params),
        unpackPrice(params),
        priceEscalationFactor(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic destination shuttle service
    case SERVICE_ITEM_CODES.DDSHUT:
      result = [
        shuttleBillableWeight(params),
        shuttleDestinationPriceDomestic(params),
        priceEscalationFactorWithoutContractYear(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic crating
    case SERVICE_ITEM_CODES.DCRT:
      result = [
        cratingSize(params, mtoParams),
        cratingPrice(params),
        priceEscalationFactorWithoutContractYear(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    // Domestic uncrating
    case SERVICE_ITEM_CODES.DUCRT:
      result = [
        cratingSize(params, mtoParams),
        unCratingPrice(params),
        priceEscalationFactorWithoutContractYear(params),
        totalAmountRequested(totalAmount),
      ];
      break;
    default:
      break;
  }
  return result;
};

export { makeCalculations as default, makeCalculations };
