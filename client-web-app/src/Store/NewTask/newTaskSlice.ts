import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface NewTaskState {
  title: string;
  description: string;
  temporary: boolean;
  date?: Date | undefined;
  weekDay?: string[] | undefined;
  time: string;
}

const initialState: NewTaskState = {
  title: '',
  description: '',
  temporary: false,
  date: undefined,
  weekDay: undefined,
  time: '00:00',
};

export const newTaskSlice = createSlice({
  name: 'newTask',
  initialState,
  reducers: {
    setTitle: (state: NewTaskState, action: PayloadAction<string>) => ({
      ...state,
      title: action.payload,
    }),
    setDescription: (state: NewTaskState, action: PayloadAction<string>) => ({
      ...state,
      description: action.payload,
    }),
    setTemporary: (state: NewTaskState, action: PayloadAction<boolean>) => ({
      ...state,
      temporary: action.payload,
    }),
    setDate: (state: NewTaskState, action: PayloadAction<Date>) => ({
      ...state,
      date: action.payload,
    }),
    setWeekDay: (state: NewTaskState, action: PayloadAction<string[]>) => ({
      ...state,
      weekDay: action.payload,
    }),
    setTime: (state: NewTaskState, action: PayloadAction<string>) => ({
      ...state,
      time: action.payload,
    }),
    resetState: () => ({
      title: '',
      description: '',
      temporary: false,
      date: undefined,
      weekDay: undefined,
      time: '00:00',
    }),
  },
});

export const {
  setTitle,
  setDescription,
  setTemporary,
  setDate,
  setWeekDay,
  setTime,
  resetState,
} = newTaskSlice.actions;

export default newTaskSlice.reducer;
