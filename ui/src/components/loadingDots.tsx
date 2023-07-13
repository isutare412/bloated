import { FC } from 'react'

export const LoadingDots: FC = () => {
  return (
    <div className="flex h-screen flex-col justify-center">
      <div className="flex justify-center">
        <span className="loading loading-dots loading-lg text-primary"></span>
      </div>
    </div>
  )
}
