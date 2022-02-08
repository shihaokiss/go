package main

import "fmt"

type AnimalIF interface {
	Eat()
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Eat() {
	fmt.Println("cat eat...")
}

func (this *Cat) GetType() (a_type string) {
	return "Cat..."
}

type Dog struct {
	color string
}

func (this *Dog) Eat() {
	fmt.Println("dog eat...")
}

func (this *Dog) GetType() (a_type string) {
	return "Dog..."
}

func Show(animal AnimalIF) {
	fmt.Println(animal.GetType())
	animal.Eat()
}

func main() {
	// cat := Cat{"write"}
	// dog := Dog{"yellow"}

	// Show(&cat)
	// Show(&dog)

	cat := &Cat{"write"}
	dog := &Dog{"yellow"}

	cat.Eat()
	dog.Eat()
	
}
