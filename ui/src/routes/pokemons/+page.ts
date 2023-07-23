import type { PageLoad } from './$types'

const defaultPokemons = [
	{ id: 702, name: '데덴네' },
	{ id: 184, name: '마릴리' },
	{ id: 777, name: '토게데마루' },
	{ id: 25, name: '피카츄' },
]

export const load = (async () => {
	return { defaultPokemons }
}) satisfies PageLoad
