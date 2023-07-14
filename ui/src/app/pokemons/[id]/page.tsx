'use client'

import { useEffect } from 'react'
import { LabeledText } from '@/components/form/labeledText'
import { LoadingDots } from '@/components/loadingDots'
import { PageTitle } from '@/components/titles'
import { pokemonVersionName } from '@/lib/model/pokemon'
import { useAppDispatch, useAppSelector } from '@/lib/redux/hooks'
import { getPokemon } from '@/lib/redux/pokemon/thunk'

export default function PokemonDetailPage({
  params,
}: {
  params: { id: number }
}) {
  const dispatch = useAppDispatch()
  const pokemonState = useAppSelector(
    (state) => state.pokemon.pokemons[params.id]
  )

  useEffect(() => {
    dispatch(getPokemon({ id: params.id }))
  }, [dispatch, params.id])

  if (
    pokemonState === undefined ||
    pokemonState.pokemon === undefined ||
    pokemonState.phase === 'pending'
  ) {
    return <LoadingDots />
  }

  const sprite = pokemonState.pokemon.spriteUrl
  const flavorTexts = pokemonState.pokemon.flavorTexts

  return (
    <div>
      <PageTitle title={pokemonState.pokemon.koreanName} />
      {sprite ? (
        // eslint-disable-next-line @next/next/no-img-element
        <img
          src={pokemonState.pokemon.spriteUrl}
          alt="pokemon sprite"
          className={`mx-auto w-1/2 max-w-xs ${
            pokemonState.pokemon.spriteUrl ? '' : 'hidden'
          }`}
        />
      ) : null}
      <div className="mt-4 flex flex-col gap-y-2">
        <LabeledText label="ID" value={pokemonState.pokemon.id.toString()} />
        <LabeledText
          label="Name"
          value={pokemonState.pokemon.name.toString()}
        />
      </div>
      {flavorTexts.length === 0 ? null : (
        <div className="mt-6">
          <label className="ml-2 text-xs font-light">Flavor Texts</label>
          <table className="mt-1 table border-t border-base-200 text-sm">
            <tbody className="border-base-content">
              {flavorTexts.map(({ version, text }) => (
                <tr key={version}>
                  <td className="font-medium">{pokemonVersionName[version]}</td>
                  <td>{text}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  )
}
