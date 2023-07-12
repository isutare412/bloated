import Link from 'next/link'
import { LikeButton } from '@/components/likeButton'
import { PageTitle } from '@/components/titles'

export default function Home() {
  return (
    <>
      <PageTitle title="Hello App Router" />
      <Link href="/streams" className="link">
        Go to streams
      </Link>
      <div className="prose max-w-none">
        <p>This is paragraph. Gonna be bloated in a while.</p>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat. Duis aute irure dolor in
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
          culpa qui officia deserunt mollit anim id est laborum.
        </p>
        <p>그리고 이건 한글이야.</p>
      </div>
      <div className="mt-4">
        <LikeButton />
      </div>
      <div className="mt-4 flex flex-wrap gap-1">
        <button className="btn">Button</button>
        <button className="btn-neutral btn">Neutral</button>
        <button className="btn-primary btn">Primary</button>
        <button className="btn-secondary btn">Secondary</button>
        <button className="btn-accent btn">Accent</button>
        <button className="btn-ghost btn">Ghost</button>
        <button className="btn-link btn">Link</button>
      </div>
    </>
  )
}
