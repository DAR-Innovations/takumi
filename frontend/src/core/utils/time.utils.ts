export const getTimeLeft = (startDate: Date, endDate: Date) => {
	const diff = endDate.getTime() - startDate.getTime()

	const months = Math.floor(diff / (1000 * 60 * 60 * 24 * 30))
	const days = Math.floor(diff / (1000 * 60 * 60 * 24)) % 30
	const hours = Math.floor((diff / (1000 * 60 * 60)) % 24)
	const minutes = Math.floor((diff / (1000 * 60)) % 60)

	let timeLeft = ''
	if (months > 0) {
		timeLeft += `${months} month${months > 1 ? 's' : ''} `
		if (days > 0) {
			timeLeft += `${days} day${days > 1 ? 's' : ''}`
		}
	} else if (days > 0) {
		timeLeft += `${days} day${days > 1 ? 's' : ''} `
		if (hours > 0) {
			timeLeft += `${hours} hour${hours > 1 ? 's' : ''}`
		}
	} else if (hours > 0) {
		timeLeft += `${hours} hour${hours > 1 ? 's' : ''} `
		if (minutes > 0) {
			timeLeft += `${minutes} minute${minutes > 1 ? 's' : ''}`
		}
	} else {
		timeLeft += `${minutes} minute${minutes > 1 ? 's' : ''}`
	}

	return timeLeft.trim()
}
