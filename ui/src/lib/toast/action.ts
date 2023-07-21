import type { ToastLevel } from '$lib/model/toast'
import { toastState } from '$lib/toast/store'

export function addToast(message: string, level: ToastLevel) {
	toastState.update((state) => {
		state.toasts.push({
			id: state.nextId++,
			message,
			level,
		})
		return state
	})
}

export function deleteToast(id: number) {
	toastState.update((state) => {
		state.toasts = state.toasts.filter((toast) => toast.id !== id)
		return state
	})
}
