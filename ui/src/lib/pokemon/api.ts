import type { GetPokemonResponse, GetPokemonSpeciesResponse } from '$lib/model/pokemon'
import { error } from '@sveltejs/kit'

export async function getPokemonSpeciesByName(
	name: string,
	options?: { fetch?: typeof fetch }
): Promise<GetPokemonSpeciesResponse> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`https://pokeapi.co/api/v2/pokemon-species/${name}`)
	if (!response.ok) {
		throw error(500, `Failed to get pokemon species by '${name}'`)
	}

	return response.json()
}

export async function getPokemonSpeciesById(
	id: number,
	options?: { fetch?: typeof fetch }
): Promise<GetPokemonSpeciesResponse> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`https://pokeapi.co/api/v2/pokemon-species/${id}`)
	if (!response.ok) {
		throw error(500, `Failed to get pokemon species by '${id}'`)
	}

	return response.json()
}

export async function getPokemonById(
	id: number,
	options?: { fetch?: typeof fetch }
): Promise<GetPokemonResponse> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`https://pokeapi.co/api/v2/pokemon/${id}`)
	if (!response.ok) {
		throw error(500, `Failed to get pokemon by '${id}'`)
	}

	return response.json()
}
