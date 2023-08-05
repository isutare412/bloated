import type { ListBannedIpsResponse } from '$lib/model/ip'
import { getBloatedApiBase } from '$lib/server/bloatedApi/base'
import { error } from '@sveltejs/kit'

export async function listBannedIps(options?: {
	fetch?: typeof fetch
}): Promise<ListBannedIpsResponse> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`${getBloatedApiBase()}/api/v1/banned-ips`)
	if (!response.ok) {
		throw error(response.status, 'Failed to list banned IPs')
	}

	return response.json()
}
