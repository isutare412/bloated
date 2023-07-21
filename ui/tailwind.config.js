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
					primary: '#4b6bfb',
					secondary: '#7b92b2',
					accent: '#67cba0',
					neutral: '#181a2a',
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
			// 		primary: '#4b6bfb',
			// 		secondary: '#7b92b2',
			// 		accent: '#67cba0',
			// 		neutral: '#181a2a',
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
