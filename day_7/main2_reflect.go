package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
	age  int
}

func (this User) Call() {
	fmt.Println("User::Call::", this)
}

func ShowUser(input interface{}) {
	input_type := reflect.TypeOf(input)
	fmt.Println("ShowUser::type::", input_type.Name())

	input_value := reflect.ValueOf(input)
	fmt.Println("ShowUser::value::", input_value)

	for i := 0; i < input_type.NumField(); i++ {
		field := input_type.Field(i)
		value := input_value.Field(i)
		fmt.Printf("name %v, type %v, value %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < input_type.NumMethod(); i++ {
		method := input_type.Method(i)
		fmt.Println("method::", method.Name, method.Type)
	}
}

func main() {
	user := User{1, "shi.hao", 20}
	user.Call()

	// 调用直接传对象
	ShowUser(user)
}
