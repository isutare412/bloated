import { FC } from 'react'

export const LabeledText: FC<{
  label: string
  value: string
  isCode?: boolean
  className?: string
}> = ({ label, value, isCode = false, className = '' }) => {
  return (
    <div className={className}>
      <label className="text-xs font-light">{label}</label>
      <span className={`block border-t ${isCode ? 'font-mono' : ''}`}>
        {value}
      </span>
    </div>
  )
}
