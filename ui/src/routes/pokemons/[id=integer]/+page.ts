import {
	pokemonVersionOrder,
	type PokemonFlavorText,
	type PokemonVersion,
} from '$lib/model/pokemon'
import { getPokemonById, getPokemonSpeciesById } from '$lib/pokemon/api'
import { error } from '@sveltejs/kit'
import { map } from 'lodash-es'
import type { PageLoad } from './$types'

export const ssr = false

type OrderedFlavorTexts = { version: PokemonVersion; text: string }[]

export const load = (async ({ params, fetch }) => {
	const id = parseInt(params.id)

	if (id < 1 || id > 1010) {
		throw error(404, `No such pokemon with ID ${id}`)
	}

	const [speciesResponse, pokemonResponse] = await Promise.all([
		getPokemonSpeciesById(id, { fetch }),
		getPokemonById(id, { fetch }),
	])

	const { names, flavor_text_entries: flavorTextEntries } = speciesResponse
	const { sprites } = pokemonResponse

	const koreanName = names.find((name) => name.language.name === 'ko')
	const flavorTexts = getOrderedFlavorTexts(flavorTextEntries)
	const spriteUrl = sprites.other['official-artwork'].front_default ?? sprites.front_default

	return {
		id,
		name: speciesResponse.name,
		koreanName: koreanName?.name,
		flavorTexts,
		spriteUrl,
	}
}) satisfies PageLoad

function getOrderedFlavorTexts(flavorTexts: PokemonFlavorText[]): OrderedFlavorTexts {
	const flavorTextByVersion: Record<string, string> = {}
	flavorTexts.forEach((flavor) => {
		if (flavor.language.name === 'en') {
			flavorTextByVersion[flavor.version.name] = flavor.flavor_text
		}
	})
	flavorTexts.forEach((flavor) => {
		if (flavor.language.name === 'ko') {
			flavorTextByVersion[flavor.version.name] = flavor.flavor_text
		}
	})

	const ordered = map(flavorTextByVersion, (text, version) => ({
		version,
		text,
	})) as OrderedFlavorTexts

	ordered.sort(({ version: left }, { version: right }) => {
		const leftIndex = pokemonVersionOrder.indexOf(left)
		const rightIndex = pokemonVersionOrder.indexOf(right)
		return leftIndex - rightIndex
	})

	return ordered
}
