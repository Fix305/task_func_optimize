package anton_m_solution

// MySuperFuncImpl MySuperFunc - реализация
// 1. `n==0` -> `x1`
// 2. `n==1` -> `x1 * x2`
// 3. `n>1` -> `f(x1, x2, n-2) * f(x1, x2, n-1)`

// var cache = make(map[uint8]float64)

func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	// cache = make(map[uint8]float64)
	// return recursive(x1, x2, n)

	if n == 0 {
		return x1
	} else if n == 1 {
		return x1 * x2
	} else {
		var preLastValue = x1
		var lastValue = x1 * x2

		var currentValue float64
		for i := uint8(2); i <= n; i++ {
			currentValue = preLastValue * lastValue

			preLastValue = lastValue
			lastValue = currentValue
		}

		return currentValue
	}
}

/*
func recursive(x1 float64, x2 float64, n uint8) float64 {
	value, exists := cache[n]

	if exists {
		return value
	} else {
		var calcValue float64
		switch n {
		case 0:
			calcValue = x1
		case 1:
			calcValue = x1 * x2
		default:
			calcValue = recursive(x1, x2, n-2) * recursive(x1, x2, n-1)
		}

		cache[n] = calcValue

		return calcValue
	}
}*/
