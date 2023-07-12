'use client'

import { animated, useSpring } from '@react-spring/web'
import { FC, useEffect, useRef } from 'react'
import { Toast } from '@/lib/model/toast'
import { useAppDispatch, useAppSelector } from '@/lib/redux/hooks'
import { deleteToast } from '@/lib/redux/toast/slice'
import { Timer } from '@/lib/utils/timer'

export const ToastContainer: FC = () => {
  const toasts = useAppSelector((state) => state.toast.toasts)

  return (
    <div className="toast-center toast w-full min-w-0 sm:w-[500px]">
      {toasts.map((toast) => (
        <ToastItem key={toast.id} {...toast} />
      ))}
    </div>
  )
}

export const ToastItem: FC<Toast> = ({ id, message, level }) => {
  const deleteTimeout = 10_000

  const dispatch = useAppDispatch()
  const timer = useRef<Timer>()

  useEffect(() => {
    timer.current = new Timer(() => {
      dispatch(deleteToast({ id, deleteAfter: 0 }))
    }, deleteTimeout)
  }, [id, dispatch])

  const [progressSpring, progressApi] = useSpring(() => {
    return {
      from: {
        value: 1,
      },
      to: {
        value: 0,
      },
      config: {
        duration: deleteTimeout,
      },
    }
  })

  const pauseTimer = () => {
    timer.current?.pause()
    progressApi.pause()
  }

  const resumeTimer = () => {
    timer.current?.resume()
    progressApi.resume()
  }

  return (
    <div
      onMouseEnter={pauseTimer}
      onMouseLeave={resumeTimer}
      className={`alert mask relative block overflow-hidden alert-${level}`}
    >
      <button
        onClick={() => dispatch(deleteToast({ id, deleteAfter: 0 }))}
        className="btn-ghost btn-sm btn-circle btn absolute right-2 top-3"
      >
        ✕
      </button>
      <div className="whitespace-pre-line break-words pr-6">{message}</div>
      <animated.progress
        className="progress absolute bottom-0 left-0 h-1 w-full bg-transparent"
        value={progressSpring.value}
      ></animated.progress>
    </div>
  )
}