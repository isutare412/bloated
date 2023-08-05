import { API_ENDPOINT } from '$env/static/private'
import type { ListBannedIpsResponse } from '$lib/model/ip'
import { error } from '@sveltejs/kit'

function getBloatedApiBase(): string {
	return API_ENDPOINT
}

export async function listBannedIps(options?: {
	fetch?: typeof fetch
}): Promise<ListBannedIpsResponse> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`${getBloatedApiBase()}/api/v1/banned-ips`)
	if (!response.ok) {
		throw error(500, 'Failed to list banned IPs')
	}

	return response.json()
}
