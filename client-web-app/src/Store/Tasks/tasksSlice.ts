import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface Task {
  id: string;
  longitude: number;
  latitude: number;
  title: string;
  description: string;
  temporary: boolean;
  date?: Date;
  weekDay?: string[];
  time: string;
}
export interface TasksState {
  data: Task[];
}

const initialState: TasksState = {
  data: [],
};

export const tasksSlice = createSlice({
  name: 'tasks',
  initialState,
  reducers: {
    setTasks: (state:TasksState, action: PayloadAction<Task[]>) => ({
      ...state,
      data: action.payload,
    }),
    addTask: (state: TasksState, action: PayloadAction<Task>) => ({
      ...state,
      data: [...state.data, action.payload],
    }),
  },
});

export const { setTasks, addTask } = tasksSlice.actions;

export default tasksSlice.reducer;
