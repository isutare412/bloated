import { BANNED_IP_FILE } from '$env/static/private'
import { watch } from 'fs'
import { readFile } from 'fs/promises'
import { forEach, once } from 'lodash-es'
import YAML from 'yaml'

const bannedIPs: Set<string> = new Set()

interface BannedIPsFile {
	ips?: {
		[country: string]: string[]
	}
}

export function isIPBanned(ip: string): boolean {
	return bannedIPs.has(ip)
}

export const keepUpdateBannedIPs = once(() => {
	console.log(`Banned IP file env: ${BANNED_IP_FILE}`)
	if (!BANNED_IP_FILE) return

	updateBannedIPs(BANNED_IP_FILE)

	watch(BANNED_IP_FILE, { persistent: false }, (_, filename) => {
		if (filename !== BANNED_IP_FILE) {
			console.error(`Unexpected banned IP file: ${filename}`)
			return
		}

		updateBannedIPs(filename)
	})
})

async function updateBannedIPs(ipsFile: string) {
	const body = await readFile(ipsFile, 'utf-8')
	const ipFile = YAML.parse(body) as BannedIPsFile | null
	if (!ipFile || !ipFile.ips) return

	bannedIPs.clear()
	forEach(ipFile.ips, (ips, country) => {
		ips.forEach((ip) => {
			bannedIPs.add(ip)
			console.log(`Now block ${country} IP: ${ip}`)
		})
	})
}
