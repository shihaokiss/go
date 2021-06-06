package main

import "fmt"

func testMap() {
    map1 := make(map[int] string)
	map1[1] = "aaa"
	map1[2] = "bbb"
	map1[3] = "ccc"
	fmt.Println(map1)

	map2 := map[string] string {
		"1": "aaa",
		"2": "bbb",
		"3": "ccc",
	}
	fmt.Println(map2)
}

func changeMap(map_test map[string]string, key string, value string) {
	map_test[key] = value
}

func mapAddDeletShow() {
    map1 := make(map[string] string)
	map1["1"] = "aaa"
	map1["2"] = "bbb"
	map1["3"] = "ccc"

	for key, value := range map1 {
		fmt.Println(key, value)
	}
    
	changeMap(map1, "1", "aaaa")
	fmt.Println("====================")

	for key, value := range map1 {
		fmt.Println(key, value)
	}
}

func main() {
    // testMap()
	mapAddDeletShow()
}