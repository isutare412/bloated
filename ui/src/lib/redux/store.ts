import { configureStore } from '@reduxjs/toolkit'
import pokemonReducer from '@/lib/redux/pokemonSlice'
import toastReducer from '@/lib/redux/toastSlice'

export const store = configureStore({
  reducer: {
    toast: toastReducer,
    pokemon: pokemonReducer,
  },
  devTools: process.env.NODE_ENV !== 'production',
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch
