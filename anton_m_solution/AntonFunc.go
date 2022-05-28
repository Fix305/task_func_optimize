package anton_m_solution

// MySuperFuncImpl MySuperFunc - реализация
// 1. `n==0` -> `x1`
// 2. `n==1` -> `x1 * x2`
// 3. `n>1` -> `f(x1, x2, n-2) * f(x1, x2, n-1)`
var cache = make(map[uint8]float64)

func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
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
			calcValue = MySuperFuncImpl(x1, x2, n-2) * MySuperFuncImpl(x1, x2, n-1)
		}

		cache[n] = calcValue

		return calcValue
	}
}