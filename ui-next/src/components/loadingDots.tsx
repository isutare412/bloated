import { FC } from 'react'

export const LoadingDots: FC = () => {
  return (
    <div className="flex h-[90vh] flex-col justify-center align-middle">
      <div className="mx-auto w-fit">
        <span className="loading loading-dots loading-lg text-primary"></span>
      </div>
    </div>
  )
}
