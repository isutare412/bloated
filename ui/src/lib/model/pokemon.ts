export interface Pokemon {
	id: number
	name: string
	koreanName: string
	flavorTexts: { version: PokemonVersion; text: string }[]
	spriteUrl?: string
}

export type LanguageType = 'ko' | 'en' | `jp-Hrkt`

export type PokemonVersion =
	| 'red'
	| 'blue'
	| 'yellow'
	| 'gold'
	| 'silver'
	| 'crystal'
	| 'ruby'
	| 'sapphire'
	| 'emerald'
	| 'firered'
	| 'leafgreen'
	| 'diamond'
	| 'pearl'
	| 'platinum'
	| 'heartgold'
	| 'soulsilver'
	| 'black'
	| 'white'
	| 'black-2'
	| 'white-2'
	| 'x'
	| 'y'
	| 'omega-ruby'
	| 'alpha-sapphire'
	| 'sun'
	| 'moon'
	| 'lets-go-pikachu'
	| 'lets-go-eevee'
	| 'ultra-sun'
	| 'ultra-moon'
	| 'sword'
	| 'shield'
	| 'legends-arceus'
	| 'scarlet'
	| 'violet'

export const pokemonVersionOrder = [
	'red',
	'blue',
	'yellow',
	'gold',
	'silver',
	'crystal',
	'ruby',
	'sapphire',
	'emerald',
	'firered',
	'leafgreen',
	'diamond',
	'pearl',
	'platinum',
	'heartgold',
	'soulsilver',
	'black',
	'white',
	'black-2',
	'white-2',
	'x',
	'y',
	'omega-ruby',
	'alpha-sapphire',
	'sun',
	'moon',
	'lets-go-pikachu',
	'lets-go-eevee',
	'ultra-sun',
	'ultra-moon',
	'sword',
	'shield',
	'legends-arceus',
	'scarlet',
	'violet',
]

export const pokemonVersionName: Record<string, string> = {
	red: 'Red',
	blue: 'Green',
	yellow: 'Yellow',
	gold: 'Gold',
	silver: 'Silver',
	crystal: 'Crystal',
	ruby: 'Ruby',
	sapphire: 'Sapphire',
	emerald: 'Emerald',
	firered: 'FireRed',
	leafgreen: 'LeafGreen',
	diamond: 'Diamond',
	pearl: 'Pearl',
	platinum: 'Platinum',
	heartgold: 'HeartGold',
	soulsilver: 'SoulSilver',
	black: 'Black',
	white: 'White',
	'black-2': 'Black 2',
	'white-2': 'White 2',
	x: 'X',
	y: 'Y',
	'omega-ruby': 'Omega Ruby',
	'alpha-sapphire': 'Alpha Sapphire',
	sun: 'Sun',
	moon: 'Moon',
	'lets-go-pikachu': "Let's Go, Picachu!",
	'lets-go-eevee': "Let's Go, Eevee!",
	'ultra-sun': 'Ultra Sun',
	'ultra-moon': 'Ultra Moon',
	sword: 'Sword',
	shield: 'Shield',
	'legends-arceus': 'Legends: Arceus',
	scarlet: 'Scarlet',
	violet: 'Violet',
}

export interface PokemonName {
	name: string
	language: {
		name: LanguageType
	}
}

export interface PokemonVariety {
	is_default: boolean
	pokemon: {
		name: string
		url: string
	}
}

export interface PokemonSprites {
	front_default?: string
	other: {
		'official-artwork': {
			front_default?: string
		}
	}
}

export interface PokemonFlavorText {
	flavor_text: string
	language: {
		name: LanguageType
	}
	version: {
		name: PokemonVersion
	}
}

export interface GetPokemonSpeciesResponse {
	id: number
	name: string
	names: PokemonName[]
	flavor_text_entries: PokemonFlavorText[]
	varieties: PokemonVariety[]
}

export interface GetPokemonResponse {
	id: number
	name: string
	sprites: PokemonSprites
}
