package main

// import "fmt" #classic style import

import ( // cleaner style
	"fmt"
)

const (
	consttest string = "this is a const"
)

var (
	vartest  string = "this is a var set in var block"
	var2test string
)

func init() {
	vartest = "this is a var set from init"
	var2test = "hi from init"
}

func main() {
	fmt.Println(consttest)
	fmt.Println(vartest)
	fmt.Println(var2test)
	var2test = "now it set from inside the function"
	fmt.Println(var2test)
}
