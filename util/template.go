package util

// 三元表达式
func TernaryExpr[T any](flag bool, a, b T) T {
	if flag {
		return a
	} else {
		return b
	}
}
