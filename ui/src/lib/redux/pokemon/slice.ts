import { createSlice } from '@reduxjs/toolkit'
import { Pokemon } from '@/lib/model/pokemon'
import { getPokemon, getPokemonDetail } from '@/lib/redux/pokemon/thunk'

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
