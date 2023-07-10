package main

import "fmt"

func Pointer() {
	a := 1993
	b := &a
	fmt.Printf("Type of b is %T\n", b)
	fmt.Println(b)
	m := *b
	fmt.Println(m)
	fmt.Println("--------------- Pointer to pointer -----------------")
	var c = 10
	var p = &c
	var pp = &p
	fmt.Println("c = ", c)
	fmt.Println("address of c = ", &c)

	fmt.Println("p = ", p)
	fmt.Println("address of p = ", &p)

	fmt.Println("pp = ", pp)

	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)
	fmt.Println("--------------- Pointer to same address -----------------")
	var d = 75
	var p1 = &d
	var p2 = &d

	if p1 == p2 {
		fmt.Println("Both pointers p1 and p2 point to the same variable.")
		fmt.Println("*p1 = ", *p1)
		fmt.Println("*p2 = ", *p2)
	}
	fmt.Println("--------------- Use pointer through func ------------------")
	x := 3
	y := 6
	px := &x
	py := &y
	fmt.Printf("Before swapping : x = %d and y = %d.\n", x, y)
	swap(px, py)
	fmt.Printf("After swapping  : x = %d and y = %d.\n", x, y)
	fmt.Println("---------------- Use pointer to struct --------------------")
	// pointer to a struct instance
	new_book := Book{120, "fiction", "Doraemon"}
	fmt.Println(new_book)
	fmt.Printf("Type of new_book -> %T\n", new_book)
	book_ptr := &new_book
	fmt.Println("Address -->", book_ptr)
	book_ptr.title = "Naruto"
	fmt.Println(new_book)
	fmt.Println("----------------- Pointer to slice --------------------")
	slsNum := [3]int{10, 12, 1993}
	fmt.Println(slsNum, "before")
	modify(slsNum[:])
	fmt.Println(slsNum, "after")
	fmt.Println("----------------- Pointer to array --------------------")
	arrNum := [3]int{3, 4, 5}
	fmt.Println(arrNum, "before")
	modifyArr(&arrNum)
	fmt.Println(arrNum, "after")
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

type Book struct {
	pages int
	genre string
	title string
}

func modify(sls []int) {
	sls[0] = 9
}
func modifyArr(arr *[3]int) {
	arr[0] = 6
}
