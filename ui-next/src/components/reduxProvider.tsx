'use client'

import React, { FC } from 'react'
import { Provider as ReduxProvider } from 'react-redux'
import { store } from '@/lib/redux/store'

export const Provider: FC<{ children: React.ReactNode }> = ({ children }) => {
  return <ReduxProvider store={store}>{children}</ReduxProvider>
}
