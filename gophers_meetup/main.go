// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// The playground is actually a 64-bit env with 32-bit pointers
// The os/arch combo is named nacl/amd64p32

// Sample program to show how to declare variables.
package main

import "fmt"
import "math"

func func1() {

	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

func func2() {

	// Declare variables that are set to their zero value.
	var a string
	var b int
	// Display the value of those variables.
	fmt.Println(a)
	fmt.Println(b)
	// Declare variables and initialize.
	// Using the short variable declaration operator.
	c := "I am c"
	d := 10
	// Display the value of those variables.
	fmt.Println(c)
	fmt.Println(d)
	// Perform a type conversion.

	// Display the value of that variable.

	fl := math.Pi
	fl2 := 3.14
	fmt.Printf("%v %T\n", fl,math.Pi)
	fmt.Println(fl)
	fmt.Printf("%v %T\n", fl2,fl2)
	fmt.Println(fl2)
}

func main() {
	func1()
	//func2()
}

/*
	Zero Values:

	Type Initialized Value
	Boolean false
	Integer 0
	Floating Point 0
	Complex 0i
	String "" (empty string)
	Pointer nil
*/