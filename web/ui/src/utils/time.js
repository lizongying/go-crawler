const isTimestampString = (str) => {
    const timestamp = Number(str);
    return !isNaN(timestamp) && timestamp >= 0;
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

export {formattedDate}