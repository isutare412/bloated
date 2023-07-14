import { createAsyncThunk } from '@reduxjs/toolkit'
import { getPokemonById, getPokemonSpeciesById } from '@/lib/api/pokemon'
import {
  Pokemon,
  PokemonVersion,
  pokemonVersionOrder,
} from '@/lib/model/pokemon'
import { AppDispatch, RootState } from '@/lib/redux/store'

export const getPokemon = createAsyncThunk<
  Pokemon,
  { id: number },
  { state: RootState; dispatch: AppDispatch }
>('pokemon/getPokemon', async ({ id }, thunkApi) => {
  var response = await getPokemonSpeciesById(id)

  const korean = response.names.find((name) => name.language.name === 'ko')
  const variety = response.varieties.find((v) => v.is_default)

  const flavorTextByVersion: Record<string, string> = {}
  response.flavor_text_entries.forEach((flavor) => {
    if (flavor.language.name === 'en') {
      flavorTextByVersion[flavor.version.name] = flavor.flavor_text
    }
  })
  response.flavor_text_entries.forEach((flavor) => {
    if (flavor.language.name === 'ko') {
      flavorTextByVersion[flavor.version.name] = flavor.flavor_text
    }
  })

  const flavorTexts = Object.keys(flavorTextByVersion).map((key) => ({
    version: key,
    text: flavorTextByVersion[key],
  })) as { version: PokemonVersion; text: string }[]

  flavorTexts.sort(({ version: left }, { version: right }) => {
    const leftIndex = pokemonVersionOrder.indexOf(left)
    const rightIndex = pokemonVersionOrder.indexOf(right)
    return leftIndex - rightIndex
  })

  if (variety) {
    thunkApi.dispatch(getPokemonDetail({ id }))
  }

  return {
    id: response.id,
    name: response.name,
    koreanName: korean!.name,
    flavorTexts,
  } satisfies Pokemon
})

export const getPokemonDetail = createAsyncThunk<
  { id: number; spriteUrl: string },
  { id: number },
  { state: RootState; dispatch: AppDispatch }
>('pokemon/getPokemonDetail', async ({ id }, thunkApi) => {
  const { sprites } = await getPokemonById(id)
  const spriteUrl =
    sprites.other['official-artwork'].front_default ?? sprites.front_default

  if (!spriteUrl) {
    throw new Error(`cannot find sprite of pokemon '${id}'`)
  }

  return { id, spriteUrl }
})
