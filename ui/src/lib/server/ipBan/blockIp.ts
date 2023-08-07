import { env } from '$env/dynamic/private'
import { listBannedIps } from '$lib/server/bloatedApi/ip'
import { once } from 'lodash-es'

const bannedIps: Set<string> = new Set()

export function isIpBanned(ip: string): boolean {
	return bannedIps.has(ip)
}

export const keepUpdateBannedIps = once(() => {
	let refreshInterval = Number(env.APP_BANNED_IP_REFRESH_SECS)
	if (isNaN(refreshInterval) || refreshInterval === 0) {
		refreshInterval = 1800
	}
	console.log(`Banned IP refresh interval: ${refreshInterval}s`)

	updateBannedIps()
	setInterval(updateBannedIps, refreshInterval * 1000)
})

async function updateBannedIps() {
	const response = await listBannedIps()
	const newBannedIps = response.bannedIps

	if (
		newBannedIps.length === bannedIps.size &&
		[...newBannedIps].every((banned) => bannedIps.has(banned.ip))
	) {
		return
	}

	bannedIps.clear()
	newBannedIps.forEach(({ country, ip }) => {
		bannedIps.add(ip)
		console.log(`Now block ${country ?? 'unknown'} IP: ${ip}`)
	})
}
