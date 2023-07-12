export interface Stream {
  name: string
  url: string
}

export const getFavoriteStreams = async (): Promise<Stream[]> => {
  await new Promise((resolve) => setTimeout(resolve, 100))
  return [
    { name: '녹두로', url: 'https://www.twitch.tv/nokduro' },
    { name: '옥냥이', url: 'https://www.twitch.tv/rooftopcat99' },
  ]
}
