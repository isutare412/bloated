import { createAsyncThunk } from '@reduxjs/toolkit'
import { AppDispatch, RootState } from '@/lib/redux/store'
import { wait } from '@/lib/utils/wait'

export const deleteToast = createAsyncThunk<
  number,
  { id: number; deleteAfter: number },
  { state: RootState; dispatch: AppDispatch }
>('toast/deleteToast', async ({ id, deleteAfter }, thunkApi) => {
  if (deleteAfter > 0) {
    await wait(deleteAfter)
  }
  return id
})
