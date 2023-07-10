package main

import (
	"errors"
	"fmt"
	"os"
)

func Errors() {
	f, err := os.Open("test.txt")
	if err != nil {
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			fmt.Println("Failed to open file at path", pErr.Path)
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
	fmt.Println("-------------- Compare Error -----------------")
	var err1 errorOne
	err2 := do()
	if err1 == err2 {
		fmt.Println("Equality Operator: Both errors are equal")
	}
	if errors.Is(err1, err2) {
		fmt.Println("Is function: Both errors are equal")
	}
}

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happended"
}
func do() error {
	return errorOne{}
}
