package utils

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Max numbers must more than 1
func Max[T Number](numbers ...T) (max T) {
	if len(numbers) == 0 {
		return
	}
	max = numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return
}

// Min numbers must more than 1
func Min[T Number](numbers ...T) (min T) {
	if len(numbers) == 0 {
		return
	}
	min = numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return
}
