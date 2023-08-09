export interface CsrfState {
	value: string
	referer: string
	createdAt: Date
}

export interface GoogleSignInRequest {
	state: string
	idToken: string
}

export interface GoogleSignInResponse {
  referer: string
}

export interface CreateCustomTokenFromGoogleIdTokenRequest {
	token: string
}

export interface CreateCustomTokenFromGoogleIdTokenResponse {
	customToken: string
}

export interface VerifyCustomTokenRequest {
	customToken: string
}

export interface CustomTokenClaims {
	email: string
	name: string
	givenName?: string
	familyName?: string
	picture?: string
}

export interface GetCustomClaimsResponse {
	claims?: CustomTokenClaims
}
