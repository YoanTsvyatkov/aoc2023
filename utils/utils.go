package utils

import (
	"os"
)

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

type MapFunc[A any, B any] func(A) B

func Map[A any, B any](input []A, m MapFunc[A, B]) []B {
	output := make([]B, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	return output
}
