const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				sans: ['Inter Variable', ...defaultTheme.fontFamily.sans],
			},
		},
	},
	plugins: [require('@tailwindcss/typography'), require('daisyui')],
	daisyui: {
		themes: [
			{
				light: {
					primary: '#f97316',
					secondary: '#fdba74',
					accent: '#84cc16',
					neutral: '#2b3440',
					'base-100': '#ffffff',
					info: '#3abff8',
					success: '#36d399',
					warning: '#fbbd23',
					error: '#f87272',
					'--btn-text-case': 'none',
				},
			},
			// {
			// 	dark: {
			// 		primary: '#f97316',
			// 		secondary: '#fdba74',
			// 		accent: '#a3e635',
			// 		neutral: '#2a323c',
			// 		'base-100': '#1d232a',
			// 		info: '#3abff8',
			// 		success: '#36d399',
			// 		warning: '#fbbd23',
			// 		error: '#f87272',
			// 		'--btn-text-case': 'none',
			// 	},
			// },
		],
	},
}
