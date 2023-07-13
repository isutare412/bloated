import {
  GetPokemonResponse,
  GetPokemonSpeciesResponse,
} from '@/lib/model/pokemon'

export async function getPokemonSpeciesByName(
  name: string
): Promise<GetPokemonSpeciesResponse> {
  const response = await fetch(
    `https://pokeapi.co/api/v2/pokemon-species/${name}`
  )
  if (!response.ok) {
    throw new Error(`failed to get pokemon species by '${name}'`)
  }

  return response.json()
}

export async function getPokemonSpeciesById(
  id: number
): Promise<GetPokemonSpeciesResponse> {
  const response = await fetch(
    `https://pokeapi.co/api/v2/pokemon-species/${id}`
  )
  if (!response.ok) {
    throw new Error(`failed to get pokemon species by '${id}'`)
  }

  return response.json()
}

export async function getPokemonByUrl(
  url: string
): Promise<GetPokemonResponse> {
  const response = await fetch(url)
  if (!response.ok) {
    throw new Error(`failed to get pokemon'`)
  }

  return response.json()
}
