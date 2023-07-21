export function debounce<T extends (...args: unknown[]) => ReturnType<T>>(
	callback: T,
	timeout = 300
): (...args: Parameters<T>) => void {
	let timer: ReturnType<typeof setTimeout>

	return (...args: Parameters<T>) => {
		clearTimeout(timer)
		timer = setTimeout(() => {
			callback(...args)
		}, timeout)
	}
}
