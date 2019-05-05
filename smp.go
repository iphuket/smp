// Package smp A Simple Mathematical Problem
package smp

// EuclideanAlgorithm 求最大公约数
func EuclideanAlgorithm(x, y int64) int64 {
	if y == 0 {
		return x
	}
	return EuclideanAlgorithm(y, x%y)
}

// DisplayProportion 最小比例
func DisplayProportion(x, y int64) (m, n int64) {
	ea := EuclideanAlgorithm(x, y)
	m, n = x/ea, y/ea
	return
}
