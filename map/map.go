package main

import "fmt"

func main() {
	amap := make(map[string]int)

	amap["key1"] = 100
	amap["key2"] = 98
	amap["key3"] = 990

	fmt.Println(amap)

	value1 := amap["key1"]
	fmt.Println("Value1 :", value1)

	fmt.Println("Length of amap is:", len(amap))

	delete(amap, "key2")
	fmt.Println(amap)

	_, isPresent := amap["key2"]
	fmt.Println("Is key2 present in the map:", isPresent)
}
