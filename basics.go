package main

// import "fmt" #classic style import

import ( // cleaner style
	"fmt"
)

const (
	constTest string = "this is a const"
)

var (
	varTest string = "this is a var set in var block"
	var2Test string
)

func init() {
	varTest = "this is a var set from init"
	var2Test = "hi from init"
}

func main() {
	fmt.Println(constTest)
	fmt.Println(varTest)
	fmt.Println(var2Test)
	var2Test = "now it set from inside the function"
	fmt.Println(var2Test)
}
