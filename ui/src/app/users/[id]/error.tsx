'use client'

import { useRouter } from 'next/navigation'
import { useEffect } from 'react'
import { useAppDispatch } from '@/lib/redux/hooks'
import { addToast } from '@/lib/redux/toastSlice'

export default function UserErrorPage({
  error,
  reset,
}: {
  error: Error
  reset: () => void
}) {
  const router = useRouter()
  const dispatch = useAppDispatch()

  useEffect(() => {
    dispatch(
      addToast({
        message: `Failed to get user: ${error.message}`,
        level: 'error',
      })
    )

    router.push('/')
  })
  return null
}
