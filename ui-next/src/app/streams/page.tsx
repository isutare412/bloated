import { PageTitle } from '@/components/titles'
import { getFavoriteStreams } from '@/lib/api/stream'

export default async function StreamsPage() {
  const favoriteStreams = await getFavoriteStreams()
  const streamItems = favoriteStreams.map((stream) => {
    return (
      <li key={stream.url}>
        Go to&nbsp;
        <a href={stream.url} className="link">
          {stream.name}
        </a>
      </li>
    )
  })

  return (
    <div>
      <PageTitle title="My Favorite Streams" />
      <ul className="list-inside list-disc">{streamItems}</ul>
    </div>
  )
}
