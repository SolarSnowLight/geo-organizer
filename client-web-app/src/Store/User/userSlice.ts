import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import axios from 'axios';

export interface UserState {
  isLoggedIn: boolean;
  access_token: string | null;
  refresh_token: string | null;
}

const initialState: UserState = {
  isLoggedIn: true,
  access_token: null,
  refresh_token: null,
};

export interface SignUpReqDTO {
  email: string;
  name: string;
  password: string;
  surname: string;
}

export interface SignInDTO {
  email: string;
  password: string;
}

export interface SignUpResDTO {
  access_token: string;
  refresh_token: string;
}

export const signUp = createAsyncThunk('user/signUp', async (user:SignUpReqDTO) => {
  const response = await axios.post<SignUpResDTO>(`${process.env.REACT_APP_API_ROUTE}/auth/sign-up/`, user);
  return response.data;
});

export const signIn = createAsyncThunk('user/signIn', async (user: SignInDTO) => {
  const response = await axios.post<SignUpResDTO>(`${process.env.REACT_APP_API_ROUTE}/auth/sign-in/`, user);
  return response.data;
});

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
  },
  extraReducers: {
    [signUp.fulfilled.name]: (state: UserState, action: PayloadAction<SignUpResDTO>) => ({
      access_token: action.payload.access_token,
      refresh_token: action.payload.refresh_token,
      isLoggedIn: true,
    }),
    [signIn.fulfilled.name]: (state: UserState, action: PayloadAction<SignUpResDTO>) => ({
      access_token: action.payload.access_token,
      refresh_token: action.payload.refresh_token,
      isLoggedIn: true,
    }),
  },
});

export default userSlice.reducer;
