import Link from 'next/link'

export default function NotFoundPage() {
  return (
    <div className="mt-[42vh] flex flex-col justify-center gap-y-6">
      <div className="flex w-full items-center justify-center font-medium">
        <span className="text-4xl">404</span>
        <div className="divider divider-horizontal h-12"></div>
        <span className="text-2xl">Not Found</span>
      </div>
      <div className="mx-auto w-fit">
        <Link href="/" className="link">
          Go Home
        </Link>
      </div>
    </div>
  )
}
