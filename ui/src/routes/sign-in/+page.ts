import type { PageLoad } from './$types'

export const load = (async ({ url }) => {
	return {
		redirectPath: url.searchParams.get('referer') ?? '/',
	}
}) satisfies PageLoad
