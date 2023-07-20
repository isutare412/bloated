import { FC } from 'react'

export const LabeledText: FC<{
  label: string
  value: string
  isCode?: boolean
  className?: string
}> = ({ label, value, isCode = false, className = '' }) => {
  return (
    <div className={className}>
      <label className="ml-2 text-xs font-light">{label}</label>
      <div className="border-t border-base-200">
        <span className={`ml-2 ${isCode ? 'font-mono' : ''}`}>{value}</span>
      </div>
    </div>
  )
}
