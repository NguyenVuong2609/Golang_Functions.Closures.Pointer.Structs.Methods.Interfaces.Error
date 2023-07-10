package main

import "fmt"

func Structs() {
	var p Person // Default all attributes
	p.Name = "Nguyễn Đức Vượng"
	p.Age = 29
	p.Nationality = "Vietnam"
	fmt.Println("Person :", p)
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Nationality:", p.Nationality)
	fmt.Println("-------------- Struct has anonymous properties --------------")
	p1 := PersonAnno{
		string: "vuong",
		int:    29,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	fmt.Println("--------------- Struct in struct -------------------")
	p2 := Person{
		Name:        "Vuong",
		Age:         29,
		Nationality: "Vietnam",
		Address: Address{
			City:     "Hanoi",
			District: "West lake",
		},
	}
	fmt.Println("Name:", p2.Name)
	fmt.Println("Age:", p2.Age)
	fmt.Println("Nationality:", p2.Nationality)
	fmt.Println("City:", p2.Address.City)
	fmt.Println("District:", p2.Address.District)
	fmt.Println("--------------- Struct equality ------------------")
	name1 := name{
		firstName: "Nguyen",
		lastName:  "Vuong",
	}
	name2 := name{
		firstName: "Nguyen",
		lastName:  "Vuong",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
}

type Person struct {
	Name        string
	Age         int
	Nationality string
	Address     Address
}
type Address struct {
	City     string
	District string
}
type PersonAnno struct {
	string
	int
}
type name struct {
	firstName string
	lastName  string
}
