package main

import "fmt"

type People struct {
	name string
	age  int
}

func (this *People) Eat() {
	fmt.Println("people eat...")
}

func (this *People) Sleep() {
	fmt.Println("people sleep...")
}

type Teacher struct {
	People
	school string
}

func (this *Teacher) Eat() {
	fmt.Println("teacher eat...")
}

func (this *Teacher) Talking() {
	fmt.Println("teacher talking...")
}

func main() {
	teacher := Teacher{People{"xxx", 20}, "aaa"}
	teacher.Eat()
	teacher.Talking()
}
