import { combineReducers } from 'redux';
import thunk from 'redux-thunk';
import { configureStore } from '@reduxjs/toolkit';
import markerReducer, { markerSlice } from './Marker/markerSlice';
import addressReducer, { addressSlice, fetchAddress } from './Address/addressSlice';
import newTaskReducer, { newTaskSlice } from './NewTask/newTaskSlice';
import appReducer, { appSlice } from './App/appSlice';
import tasksReducer, { fetchTasks, tasksSlice } from './Tasks/tasksSlice';
import userReducer, { signIn, signUp } from './User/userSlice';

const addressActions = addressSlice.actions;
const markerActions = markerSlice.actions;
const newTaskActions = newTaskSlice.actions;
const appActions = appSlice.actions;
const tasksActions = tasksSlice.actions;

export const ActionCreators = {
  ...addressActions,
  ...markerActions,
  ...newTaskActions,
  ...appActions,
  ...tasksActions,
  fetchAddress,
  signUp,
  signIn,
  fetchTasks,
};

const rootReducer = combineReducers({
  marker: markerReducer,
  address: addressReducer,
  newTask: newTaskReducer,
  app: appReducer,
  tasks: tasksReducer,
  user: userReducer,
});

export const store = configureStore({
  reducer: rootReducer,
  middleware: [thunk],
  devTools: true,
});
export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
