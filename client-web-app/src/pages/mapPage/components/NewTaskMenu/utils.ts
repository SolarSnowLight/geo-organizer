import { GeocodeType } from '../../../../types/GeocodeTypes';

export const getAddress = (data: GeocodeType | undefined) : string => {
  const address: string = typeof data !== 'undefined'
    ? `${data && data.features[0].context[1].text}, ${
      data.features[0].text
    }, д.${data.features[0].address}`
    : 'Поставьте метку на карту';
  return address;
};

export default getAddress;
