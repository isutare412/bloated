import { env } from '$env/dynamic/private'

export function getBloatedApiBase(): string {
	return env.APP_API_ENDPOINT
}
