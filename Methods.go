package main

import "fmt"

func Methods() {
	p := Rectangle{2.0, 4.0}
	s := Square{4}

	fmt.Println("Point Rectangle: ", p)
	fmt.Println("Rectangle Area is : ", p.Acreage())
	fmt.Println("Point Square: ", s)
	fmt.Println("Square Area is : ", s.Acreage())
	fmt.Println("-------------- Receiver Pointer ----------------")
	t := Point{3, 4}
	fmt.Println("Point t = ", t)

	t.Translate(7, 9)
	fmt.Println("After Translate, t = ", t)
	fmt.Println("--------------- Method non-struct ---------------")
	myStr := MyString("OLLEH")
	fmt.Println(myStr, "<-- Before")
	fmt.Println(myStr.reverse(), "<-- After")
}

type MyString string

func (myStr MyString) reverse() string {
	s := string(myStr)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type Point struct {
	X, Y float64
}

func (p *Point) Translate(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// Struct type - `Rectangle`
type Rectangle struct {
	X, Y float64
}
type Square struct {
	X float64
}

// Method with receiver `Square`
func (s Square) Acreage() float64 {
	return s.X * s.X
}

// Method with receiver `Rectangle`
func (p Rectangle) Acreage() float64 {
	return p.Y * p.X
}
