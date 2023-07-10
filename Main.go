package main

import "fmt"

func main() {
	fmt.Println("------------ Function ------------")
	fmt.Println(Add(2, 3))
	fmt.Println(Plus(4, 5))
	fmt.Println("------------ Multi return values ---------")
	fmt.Println(calculate(5, 8))
	fmt.Println("------------ Blank Identifier -------------")
	area, _ := calculate(10, 5)
	fmt.Printf("Area %d ", area)
	fmt.Println("------------ slice parameter --------------")
	numSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(0, numSlice...))
	fmt.Println("------------ Defer -------------")
	fmt.Println("hello")
	fmt.Println(yes())
	fmt.Println("------------------- Closure ------------------------")
	Closures()
	fmt.Println("------------------- Factorial - Closure --------------------")
	f := Factorial()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println("------------------- Pointer ------------------------")
	Pointer()
	fmt.Println("------------------- Structs ------------------------")
	Structs()
	fmt.Println("------------------- Methods ------------------------")
	Methods()
	fmt.Println("------------------- Interfaces ---------------------")
	Interfaces()
	fmt.Println("------------------- Errors -------------------------")
	Errors()
}

// Add function named
func Add(a, b int) int {
	return a + b
}

// Plus anonymous function
var Plus = func(a, b int) int {
	return a + b
}

// Multi return values, named return values
func calculate(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = (length + width) * 2
	return
}
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

func yes() (text string) {
	defer func() {
		text = "no"
	}()
	return "yes"
}
