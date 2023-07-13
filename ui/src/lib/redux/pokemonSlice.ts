import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { getPokemonByUrl, getPokemonSpeciesById } from '@/lib/api/pokemon'
import {
  Pokemon,
  PokemonVersion,
  pokemonVersionOrder,
} from '@/lib/model/pokemon'
import { AppDispatch, RootState } from '@/lib/redux/store'

export interface PokemonState {
  phase: 'pending' | 'fetched' | 'failed'
  pokemon?: Pokemon
}

export interface PokemonsState {
  pokemons: Partial<Record<number, PokemonState>>
}

const initialState: PokemonsState = {
  pokemons: {},
}

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
    thunkApi.dispatch(getPokemonDetail({ url: variety.pokemon.url }))
  }

  return {
    id: response.id,
    name: response.name,
    koreanName: korean!.name,
    flavorTexts,
  } satisfies Pokemon
})

const getPokemonDetail = createAsyncThunk<
  { id: number; spriteUrl: string },
  { url: string },
  { state: RootState; dispatch: AppDispatch }
>('pokemon/getPokemonDetail', async ({ url }, thunkApi) => {
  const { id, sprites } = await getPokemonByUrl(url)
  const spriteUrl =
    sprites.other['official-artwork'].front_default ?? sprites.front_default

  if (!spriteUrl) {
    throw new Error(`cannot find sprite of pokemon '${id}'`)
  }

  return { id, spriteUrl }
})

export const pokemonSlice = createSlice({
  name: 'pokemon',
  initialState: initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(getPokemon.pending, (state, action) => {
      const pokeId = action.meta.arg.id

      const pokeState = state.pokemons[pokeId]
      if (pokeState === undefined) {
        state.pokemons[pokeId] = { phase: 'pending' }
      } else {
        pokeState.phase = 'pending'
      }
    })

    builder.addCase(getPokemon.rejected, (state, action) => {
      const pokdId = action.meta.arg.id

      const pokeState = state.pokemons[pokdId]
      if (pokeState === undefined) {
        state.pokemons[pokdId] = { phase: 'failed' }
      } else {
        pokeState.phase = 'failed'
      }
    })

    builder.addCase(getPokemon.fulfilled, (state, action) => {
      const pokeId = action.meta.arg.id

      state.pokemons[pokeId] = {
        pokemon: action.payload,
        phase: 'fetched',
      }
    })

    builder.addCase(getPokemonDetail.fulfilled, (state, action) => {
      const pokeId = action.payload.id

      const pokemon = state.pokemons[pokeId]?.pokemon
      if (pokemon) {
        pokemon.spriteUrl = action.payload.spriteUrl
      }
    })
  },
})

export default pokemonSlice.reducer
