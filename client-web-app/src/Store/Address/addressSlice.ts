import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import axios from 'axios';
import { GeocodeType } from '../../types/GeocodeTypes';
import { MarkerType } from '../../pages/mapPage/components/CustomMap';
import { addMarker } from '../Marker/markerSlice';

export interface AddressState {
  data: GeocodeType | undefined;
}

const initialState: AddressState = {
  data: undefined,
};
const token = 'pk.eyJ1IjoiaW5ub2tlbnRpeTI1MTciLCJhIjoiY2wwdHRicHd2MHAxZjNibm1odTdwNXk1cCJ9.aibsBxys2tKJkN25qkCAKg';

export const addressSlice = createSlice({
  name: 'address',
  initialState,
  reducers: {
    addAddress: (state, action: PayloadAction<GeocodeType>) => ({
      ...state,
      data: action.payload,
    }),
    deleteAddress: (state) => ({
      ...state,
      data: undefined,
    }),
  },
});

export const {
  addAddress,
  deleteAddress,
} = addressSlice.actions;

export const fetchAddress = createAsyncThunk('address/fetchAddress', async (latlng:MarkerType, { dispatch }) => {
  dispatch(addMarker(latlng));
  try {
    const response = await axios.get<GeocodeType>(
      `https://api.mapbox.com/geocoding/v5/mapbox.places/${latlng.longitude},${latlng.latitude}.json?access_token=${token}`,
    );
    dispatch(addAddress(response.data));
  } catch (e) {
    console.log(e);
  }
});

export default addressSlice.reducer;
