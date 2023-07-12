import 'client-only'
import { PayloadAction, createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { Toast, ToastLevel } from '@/lib/model/toast'

export interface ToastState {
  toasts: Toast[]
}

const initialState: ToastState = {
  toasts: [],
}

let nextToastId = 0

export const deleteToast = createAsyncThunk(
  'toast/deleteToast',
  async (
    { id, deleteAfter }: { id: number; deleteAfter: number },
    thunkAPI
  ) => {
    await new Promise((resolve) => setTimeout(resolve, deleteAfter))
    return id
  }
)

export const toastSlice = createSlice({
  name: 'toast',
  initialState: initialState,
  reducers: {
    addToast: (
      state,
      action: PayloadAction<{ message: string; level: ToastLevel }>
    ) => {
      state.toasts.push({
        id: nextToastId++,
        message: action.payload.message,
        level: action.payload.level,
      })
    },
  },
  extraReducers: (builder) => {
    builder.addCase(deleteToast.fulfilled, (state, action) => {
      const toastId = action.payload
      state.toasts = state.toasts.filter((toast) => toast.id !== toastId)
    })
  },
})

export const toastReducer = toastSlice.reducer

export const addToast = toastSlice.actions.addToast
