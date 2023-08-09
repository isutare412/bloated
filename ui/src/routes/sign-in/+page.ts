import type { PageLoad } from './$types'

export const load = (async (event) => {
	const { url } = event

	return {
		redirectPath: url.searchParams.get('referer') ?? '/',
	}
}) satisfies PageLoad
