import type { Toast } from '$lib/model/toast'
import { writable } from 'svelte/store'

interface ToastState {
	nextId: number
	toasts: Toast[]
}

export const toastState = writable<ToastState>({ nextId: 0, toasts: [] })
