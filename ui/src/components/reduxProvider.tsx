'use client'

import React, { FC } from 'react'
import { Provider } from 'react-redux'
import { store } from '@/lib/redux/store'

export const ReduxProvider: FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  return <Provider store={store}>{children}</Provider>
}
