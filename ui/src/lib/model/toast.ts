export type ToastLevel = 'info' | 'success' | 'warning' | 'error'

export interface Toast {
	id: number
	message: string
	level: ToastLevel
}
