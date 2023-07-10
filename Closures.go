package main

import "fmt"

func Closures() {
	c1 := incrementer()

	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	c2 := incrementer()

	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())

}
func incrementer() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}
func Factorial() func() int {
	fact, n := 1, 1
	return func() int {
		fact = fact * n
		n += 1
		return fact
	}
}
