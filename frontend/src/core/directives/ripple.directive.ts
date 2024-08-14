// src/directives/ripple.ts

import type { Directive } from 'vue'

function createRipple(event: MouseEvent): void {
	const button = event.currentTarget as HTMLElement

	const circle = document.createElement('span')
	const diameter = Math.max(button.clientWidth, button.clientHeight)
	const radius = diameter / 2

	circle.style.width = circle.style.height = `${diameter}px`
	circle.style.left = `${event.clientX - button.offsetLeft - radius}px`
	circle.style.top = `${event.clientY - button.offsetTop - radius}px`
	circle.classList.add('ripple')

	const ripple = button.getElementsByClassName('ripple')[0]

	if (ripple) {
		ripple.remove()
	}

	button.appendChild(circle)
}

export const rippleDirective: Directive = {
	mounted(el: HTMLElement) {
		el.classList.add('relative', 'overflow-hidden')
		el.addEventListener('click', createRipple)
	},
	unmounted(el: HTMLElement) {
		el.removeEventListener('click', createRipple)
	},
}
