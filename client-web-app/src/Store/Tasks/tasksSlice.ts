import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import mocktasks from './mocktasks';

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
  completed: boolean;
  address: string;
}
export interface TasksState {
  data: Task[];
}

const initialState: TasksState = {
  data: [],
};

// // eslint-disable-next-line no-unused-vars
// export const fetchTasks = createAsyncThunk('tasks/fetchTasks', (_, { dispatch }) => {
//   // eslint-disable-next-line no-use-before-define
//   dispatch(setTasks(mocktasks));
// });

export const fetchTasks = createAsyncThunk('tasks/fetchTasks', () => mocktasks);

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
  extraReducers: {
    /* builder.addCase(fetchTasks.fulfilled, (state, action) => ({
      ...state,
      data: action.payload,
    })); */
    [fetchTasks.fulfilled.type]: (state, action: PayloadAction<Task[]>) => ({
      ...state,
      data: action.payload,
    }),
  },
});

export const { setTasks, addTask } = tasksSlice.actions;

export default tasksSlice.reducer;
