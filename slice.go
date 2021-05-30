package main

import "fmt"

func main() {
	var myArray = [...]int{1, 2, 3}
	var mySlice = myArray[:]
	mySlice = append(mySlice,100)
	fmt.Println(myArray)
	fmt.Println(mySlice)

	var mySlice2 = []float32{14.,15.,19.}
	fmt.Println(mySlice2)

	var mySlice3 = make([]float32,100)
	mySlice3[0]= 12.
	mySlice3[1] = 13.
	mySlice3[2] = 14.
	fmt.Println(mySlice3)
}
