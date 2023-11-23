const isTimestampString = (str) => {
    const timestamp = Number(str);
    return !isNaN(timestamp) && timestamp > 0;
}

const formattedDate = (timestamp) => {
    if (!isTimestampString(timestamp)) {
        return ''
    }
    const date = new Date(timestamp * 1000)
    const year = date.getFullYear()
    const month = (date.getMonth() + 1).toString().padStart(2, '0')
    const day = date.getDate().toString().padStart(2, '0')
    const hour = date.getHours().toString().padStart(2, '0')
    const minute = date.getMinutes().toString().padStart(2, '0')
    const second = date.getSeconds().toString().padStart(2, '0')
    return `${year}-${month}-${day} ${hour}:${minute}:${second}`
}

function formatDuration(seconds) {
    if (seconds < 0) {
        seconds = Math.round(new Date().getTime() / 1000) + seconds
    }
    if (seconds < 60) {
        return seconds === 1 ? '1 second' : `${seconds} seconds`
    } else if (seconds < 3600) {
        const minutes = Math.floor(seconds / 60);
        return minutes === 1 ? '1 minute' : `${minutes} minutes`
    } else if (seconds < 86400) {
        const hours = Math.floor(seconds / 3600);
        return hours === 1 ? '1 hour' : `${hours} hours`
    } else {
        const days = Math.floor(seconds / 86400);
        return days === 1 ? '1 day' : `${days} days`
    }
}

export {formattedDate, formatDuration}