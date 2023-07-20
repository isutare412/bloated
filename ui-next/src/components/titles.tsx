import { FC } from 'react'

export const PageTitle: FC<{ title: string; className?: string }> = ({
  title,
  className = '',
}) => {
  return (
    <h1 className={`mb-6 mt-3 text-4xl font-semibold ${className}`}>{title}</h1>
  )
}
