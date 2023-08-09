import type {
	CreateCustomTokenFromGoogleIdTokenRequest,
	CreateCustomTokenFromGoogleIdTokenResponse,
	CustomTokenClaims,
	VerifyCustomTokenRequest,
} from '$lib/model/auth'
import { getBloatedApiBase } from '$lib/server/bloatedApi/base'
import { error } from '@sveltejs/kit'

export async function createCustomTokenFromGoogleIdToken(
	request: CreateCustomTokenFromGoogleIdTokenRequest,
	options?: { fetch?: typeof fetch }
): Promise<CreateCustomTokenFromGoogleIdTokenResponse> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`${getBloatedApiBase()}/api/v1/tokens?issuer=google`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(request),
	})
	if (!response.ok) {
		throw error(response.status, `Failed to exchange custom token from Google ID token`)
	}

	return response.json()
}

export async function verifyCustomToken(
	customToken: string,
	options?: { fetch?: typeof fetch }
): Promise<CustomTokenClaims> {
	const customFetch = options?.fetch ?? fetch

	const request = {
		customToken,
	} satisfies VerifyCustomTokenRequest

	const response = await customFetch(`${getBloatedApiBase()}/api/v1/tokens/verify`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(request),
	})
	if (!response.ok) {
		throw error(response.status, `Failed to verify custom token`)
	}

	return response.json()
}
