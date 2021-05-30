package main

const (
	first = iota // auto increment
	second
)

const(
	third = iota // reset to 0 with new const block
)

func main() {
	println(first)
	println(second)
	println(third)
}
