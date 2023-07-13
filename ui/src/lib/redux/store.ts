import { configureStore } from '@reduxjs/toolkit'
import toastReducer from '@/lib/redux/toastSlice'

export const store = configureStore({
  reducer: {
    toast: toastReducer,
  },
  devTools: process.env.NODE_ENV !== 'production',
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch
