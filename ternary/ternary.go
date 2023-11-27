package ternary

func Ternary[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	}
	return ifFalse
}
