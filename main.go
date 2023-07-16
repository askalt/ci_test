package main

import (
	"fmt"
	"main/lib"
	"os"
	"strconv"
)

func getArg(args []string, i int) (int, error) {
	if len(args) < i+1 {
		return 0, nil
	}
	return strconv.Atoi(args[i])
}

func main() {
	if len(os.Args) < 2 {
		panic("provide aciton")
	}
	a, err := getArg(os.Args, 2)
	if err != nil {
		panic(err)
	}
	b, err := getArg(os.Args, 3)
	if err != nil {
		panic(err)
	}
	result, err := lib.Calc(os.Args[1], a, b)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
