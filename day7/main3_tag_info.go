package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	name string `info:"学生名称" doc:"年龄小于20岁"`
	age  int    `info:"性别"`
}

func ShowTags(user interface{}) {
	user_type := reflect.TypeOf(user).Elem()
	for i := 0; i < user_type.NumField(); i++ {
		info_tag := user_type.Field(i).Tag.Get("info")
		doc_tag := user_type.Field(i).Tag.Get("doc")
		fmt.Println(user_type.Field(i).Name, user_type.Field(i).Type, info_tag, doc_tag)
	}
}

func test() {
	var user1 User
	ShowTags(&user1)
}

////////////////////////////////////////////////////////
type Movice struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Money  int    `json:"money"`
	People []string
}

func test2() {
	// json 序列化
	movice := Movice{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}
	jsonStr, err := json.Marshal(movice)
	if err != nil {
		fmt.Println("json.Marshal::", err)
		return
	}
	fmt.Printf("movice_info::%s\n", jsonStr)

	// json 反序列化
	movice2 := Movice{}
	err = json.Unmarshal(jsonStr, &movice2)
	if err != nil {
		fmt.Println("json.unmarshal::", err)
		return
	}
	fmt.Println(movice2)
}

func main() {
	test2()
}
