package main

import (
	"fmt"
	"strconv"
)

func Interfaces() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	// define interface i
	var i Men

	//i can store Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i can store Employee
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// these three elements are different types but they all implemented interface Men
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
	fmt.Println("----------------- Use pointer interface ------------------")
	player := Player{name: "Steve", health: 100}
	mob := Mob{name: "Zombie", health: 140}
	fmt.Println(player.intro())
	player.SayHi()
	fmt.Println(mob.intro())
	fmt.Println(mob)
	fmt.Println(player)
	fmt.Println(player.attack(&mob.health))
	fmt.Println(mob.attack(&player.health))
	fmt.Println(mob)
	fmt.Println(player)
	fmt.Println("------------- Void interface ----------------")
	// vars
	u := 5
	v := "Hello world"

	// lưu lại già trị
	void = u
	fmt.Println(u, "<--- u")
	void = v
	fmt.Println(v, "<--- v")
	fmt.Println("------------- Void interface -----------------")
	num := parseInt(4)
	fmt.Println(num)
	num = parseInt("4")
	fmt.Println(num)
	num = parseInt(4.1243)
	fmt.Println(num)
}

var void interface{}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// Interface Men implemented by Human, Student and Employee
type Men interface {
	SayHi()
	Sing(lyrics string)
}

// method
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// method
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// method
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

// Use pointer
type Creature interface {
	intro() string
	attack(*int) int
	Men
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) SayHi() {
	fmt.Printf("Hi, I am %s my health is %d\n", p.name, p.health)
}

func (p Player) intro() string {
	fmt.Println("Player has spawned")
	return p.name
}

func (p Player) attack(m_health *int) int {
	fmt.Println("Player has attacked!")
	*m_health = *m_health - 50
	return *m_health
}

func (m Mob) intro() string {
	fmt.Printf("A wild %s has appeared!\n", m.name)
	return m.name
}
func (m Mob) attack(p_health *int) int {
	fmt.Printf("%s has attacked you! -%d\n", m.name, 30)
	*p_health = *p_health - 30
	return *p_health
}

// Void interfaces
func parseInt(n interface{}) int {
	switch n.(type) {
	case int:
		return (n).(int) * (n).(int)
	case string:
		s, _ := strconv.Atoi(n.(string))
		return s
	case float64:
		return int(n.(float64))
	default:
		return n.(int)
	}
}
