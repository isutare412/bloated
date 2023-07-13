'use client'

import { FC, useState } from 'react'
import { useAppDispatch } from '@/lib/redux/hooks'
import { addToast } from '@/lib/redux/toastSlice'

export const ToastEmitter: FC = () => {
  const dispatch = useAppDispatch()

  const [message, setMessage] = useState('Burn some toasts :)')

  return (
    <div>
      <p className="text-lg">Toast Tester</p>
      <div>
        <label className="text-xs font-light">
          <span className="label-text">Toast message</span>
        </label>
        <input
          type="text"
          placeholder="Type here"
          onChange={(event) => setMessage(event.target.value)}
          value={message}
          className="input-bordered input input-sm block w-full max-w-lg"
        />
      </div>
      <div className="mt-2 flex flex-wrap gap-2">
        <button
          onClick={() => dispatch(addToast({ message, level: 'info' }))}
          className="btn-info btn"
        >
          Info
        </button>
        <button
          onClick={() => dispatch(addToast({ message, level: 'success' }))}
          className="btn-success btn"
        >
          Success
        </button>
        <button
          onClick={() => dispatch(addToast({ message, level: 'warning' }))}
          className="btn-warning btn"
        >
          Warning
        </button>
        <button
          onClick={() => dispatch(addToast({ message, level: 'error' }))}
          className="btn-error btn"
        >
          Error
        </button>
      </div>
    </div>
  )
}
