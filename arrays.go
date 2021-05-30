package main

import "fmt"

func main() {
	var myArray = [...]string{"1", "2", "3"}
	fmt.Println(myArray)

	var myArray2 [2]string
	myArray2[0] = "Hello"
	myArray2[1] = "World"
	fmt.Println(myArray2[0], myArray2[1])
	fmt.Println(myArray2)

	myArray3 := [3]int{}
	myArray3[0] = 42
	myArray3[1] = 27
	myArray3[2] = 99

	fmt.Println(myArray3)
	fmt.Println(len(myArray))
}
