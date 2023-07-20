'use client'

import Link from 'next/link'
import { useEffect } from 'react'
import { useAppDispatch } from '@/lib/redux/hooks'
import { addToast } from '@/lib/redux/toast/slice'

export default function ErrorPage({
  error,
  reset,
}: {
  error: Error
  reset: () => void
}) {
  const dispatch = useAppDispatch()

  useEffect(() => {
    dispatch(
      addToast({
        message: error.message,
        level: 'error',
      })
    )
  })

  return (
    <div className="mt-[42vh] flex flex-col justify-center gap-y-6">
      <div className="mx-auto w-fit">
        <span className="text-2xl">Unexpected Error</span>
      </div>
      <div className="mx-auto w-fit">
        <Link href="/" className="link">
          Go Home
        </Link>
      </div>
    </div>
  )
}
