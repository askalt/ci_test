package lib

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func Calc(typ string, a, b int) (int, error) {
	if typ == "add" {
		return add(a, b), nil
	} else if typ == "mul" {
		return mul(a, b), nil
	}
	return 0, fmt.Errorf("unexpected action")
}
