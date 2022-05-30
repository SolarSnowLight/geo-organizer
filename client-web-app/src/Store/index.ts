import { combineReducers } from 'redux';
import thunk from 'redux-thunk';
import { configureStore } from '@reduxjs/toolkit';
import markerReducer, { markerSlice } from './Marker/markerSlice';
import addressReducer, { addressSlice, fetchAddress } from './Address/addressSlice';
import newTaskReducer, { newTaskSlice } from './NewTask/newTaskSlice';

const addressActions = addressSlice.actions;
const markerActions = markerSlice.actions;
const newTaskActions = newTaskSlice.actions;

export const ActionCreators = {
  ...addressActions,
  ...markerActions,
  ...newTaskActions,
  fetchAddress,
};

const rootReducer = combineReducers({
  marker: markerReducer,
  address: addressReducer,
  newTask: newTaskReducer,
});

export const store = configureStore({
  reducer: rootReducer,
  middleware: [thunk],
  devTools: true,
});
export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
