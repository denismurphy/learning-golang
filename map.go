package main

import "fmt"

func main(){
	var myMap = make(map[int]string)
	myMap[1] = "test"
	myMap[2] = "hello"
	fmt.Println(myMap)
}
