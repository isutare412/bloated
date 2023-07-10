import { LikeButton } from '@/components/likeButton'

export default function Home() {
  return (
    <div className="mx-auto p-2 md:max-w-3xl">
      <div className="prose max-w-none">
        <h1>Hello Next.JS!</h1>
        <p>This is paragraph. Gonna be bloated in a while.</p>
      </div>
      <div className="mt-4">
        <LikeButton />
      </div>
    </div>
  )
}
