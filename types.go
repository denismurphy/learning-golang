package main

func main() {
	var myInt int
	myInt = 42
	println(myInt)

	var myFloat32 float32 = 42. //dot after number make it a float literal
	println(myFloat32)

	myString := "string" // short hand type inference
	println(myString)

	var myComplex complex64 = complex(1,2)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))
}
