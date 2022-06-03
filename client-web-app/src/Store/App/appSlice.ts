import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export enum BarEnum {
  NEW_TASK = 'NEW_TASK',
  TASK_LIST = 'TASK_LIST',
}

export enum TabEnum {
  CURRENT = 'CURRENT',
  PERSISTENT = 'PERSISTENT',
  ACCOMPLISHED = 'ACCOMPLISHED',
}

export interface AppState {
  bar: BarEnum;
  tab: TabEnum;
}

const initialState: AppState = {
  bar: BarEnum.TASK_LIST,
  tab: TabEnum.CURRENT,
};

export const appSlice = createSlice({
  name: 'app',
  initialState,
  reducers: {
    setBar: (state, action: PayloadAction<BarEnum>) => ({
      ...state,
      bar: action.payload,
    }),
    setTab: (state, action: PayloadAction<TabEnum>) => ({
      ...state,
      tab: action.payload,
    }),
  },
});

export const { setTab, setBar } = appSlice.actions;

export default appSlice.reducer;
