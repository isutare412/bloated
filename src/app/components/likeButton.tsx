'use client'

import { FC, useState } from 'react'

export const LikeButton: FC<{ className?: string }> = ({ className = '' }) => {
  const [likeCount, setLikeCount] = useState(0)

  function handleClick() {
    setLikeCount(likeCount + 1)
  }

  return (
    <button onClick={handleClick} className={`btn ${className}`}>
      Like ({likeCount})
    </button>
  )
}
