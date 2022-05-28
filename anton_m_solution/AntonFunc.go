package anton_m_solution

import "math"

// MySuperFuncImpl MySuperFunc - реализация
// 1. `n==0` -> `x1`
// 2. `n==1` -> `x1 * x2`
// 3. `n>1` -> `f(x1, x2, n-2) * f(x1, x2, n-1)`

var cache = make(map[uint8]float64)

func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	cache = make(map[uint8]float64)

	// Более шустрый вариант, но 2 теста не проходит по точности вычисления
	/// return flat(x1, x2, n)
	return recursive(x1, x2, n)
}

// Более шустрый вариант, но 2 теста не проходит по точности вычисления
func flat(x1 float64, x2 float64, n uint8) float64 {
	var lastX1coef uint8 = 1
	var lastX2coef uint8 = 0
	var currentX1coef uint8 = 0
	var currentX2coef uint8 = 0

	for i := uint8(1); i <= n; i++ {
		currentX1coef = lastX1coef + lastX2coef
		currentX2coef = lastX1coef

		lastX1coef = currentX1coef
		lastX2coef = currentX2coef
	}

	if n == 0 {
		return math.Pow(x1, float64(lastX1coef))
	} else {
		return math.Pow(x1, float64(lastX1coef)) * math.Pow(x2, float64(lastX2coef))
	}
}

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
}
