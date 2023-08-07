import { isIpBanned } from '$lib/server/ipBan/blockIp'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load = (async ({ getClientAddress }) => {
	const ip = getClientAddress()
	console.log(`IP: ${ip}`)

	if (isIpBanned(ip)) {
		console.log(`Block banned IP: ${ip}`)
		throw error(403, 'Banned IP')
	}
}) satisfies LayoutServerLoad
