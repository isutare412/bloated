'use client'

import { useRouter } from 'next/navigation'
import { useEffect } from 'react'
import { LabeledText } from '@/components/form/labeledText'
import { LoadingDots } from '@/components/loadingDots'
import { PageTitle } from '@/components/titles'
import { pokemonVersionName } from '@/lib/model/pokemon'
import { useAppDispatch, useAppSelector } from '@/lib/redux/hooks'
import { getPokemon } from '@/lib/redux/pokemon/thunk'
import { addToast } from '@/lib/redux/toast/slice'

export default function PokemonDetailPage({
  params,
}: {
  params: { id: number }
}) {
  const router = useRouter()
  const dispatch = useAppDispatch()
  const pokemonState = useAppSelector(
    (state) => state.pokemon.pokemons[params.id]
  )

  useEffect(() => {
    dispatch(getPokemon({ id: params.id }))
  }, [dispatch, params.id])

  useEffect(() => {
    if (pokemonState?.phase === 'failed') {
      dispatch(
        addToast({
          message: `Failed to catch pokemon of ID '${params.id}'`,
          level: 'error',
        })
      )
      router.push('/')
    }
  }, [router, dispatch, params.id, pokemonState?.phase])

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
          <label className="text-lg font-light">Flavor Texts</label>
          <table className="mt-2 table border-t border-base-200 text-xs">
            <tbody className="border-base-content">
              {flavorTexts.map(({ version, text }) => (
                <tr key={version}>
                  <td>{pokemonVersionName[version]}</td>
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
