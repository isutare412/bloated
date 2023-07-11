import { FC } from 'react'

export const PageTitle: FC<{ title: string }> = ({ title }) => {
  return <h1 className="mb-6 text-4xl font-semibold text-primary">{title}</h1>
}
