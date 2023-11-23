const sortInt = (a, b) => {
    if (a === b) {
        return 0
    }
    const a1 = a === undefined ? 0 : a
    const b1 = b === undefined ? 0 : b
    return a1 - b1
}

const sortBigInt16 = (a, b) => {
    if (a === b) {
        return 0
    }
    const a1 = a === undefined ? '0' : a.slice(0, 20)
    const b1 = b === undefined ? '0' : b.slice(0, 20)
    return BigInt(`0x${a1}`) - BigInt(`0x${b1}`) > 0 ? 1 : -1
}

const sortBigInt = (a, b) => {
    if (a === b) {
        return 0
    }
    const a1 = a === undefined ? '0' : a
    const b1 = b === undefined ? '0' : b
    return BigInt(a1) - BigInt(b1) > 0 ? 1 : -1
}

const sortStr = (a, b) => {
    if (a === b) {
        return 0
    }
    const a1 = a === undefined ? '' : a
    const b1 = b === undefined ? '' : b
    return a1.localeCompare(b1)
}

export {sortInt, sortBigInt, sortBigInt16, sortStr}