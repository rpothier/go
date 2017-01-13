// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// The playground is actually a 64-bit env with 32-bit pointers
// The os/arch combo is named nacl/amd64p32

// Sample program to show how to declare variables.
package main

import (
	"fmt"
	"github.com/rpothier/go/gophers_meetup/exercises"
	"github.com/rpothier/go/gophers_meetup/toy"
	"math"
)

func var_ex1() {

	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int     %T [%v]\n", a, a)
	fmt.Printf("var b string  %T [%v]\n", b, b)
	fmt.Printf("var c float64 %T [%v]\n", c, c)
	fmt.Printf("var d bool    %T [%v]\n\n", d, d)

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

	// Perform a type conversion
	// Display the value of that variable
	//var pi float32

	pi := float32(math.Pi)
	fmt.Printf("%v %T\n", math.Pi, math.Pi)
	fmt.Printf("%v %T\n", pi, pi)
}

func struct_ex1() {

	type user struct {
		name  string
		email string
		age   int
	}
	u := user{"Fred", "Fred@email.com", 32}
	fmt.Println(u.name, u.email, u.age)

	u2 := struct {
		name  string
		email string
		age   int
	}{
		"ethel", "ethel@email.com", 30,
	}
	fmt.Println(u2.name, u2.email, u2.age)

}

func pointer_ex1() {
	var i int
	i = 20
	fmt.Printf("i = %v addr = %v\n", i, &i)
	p := &i
	fmt.Printf("addr of p = %v , value of p =%v p points to %v\n", &p, p, *p)
	pointer_ex2()
}

func pointer_ex2() {
	type xt struct {
		name string
		age  int
	}

	x := xt{"Suzy", 21}

	//y := func (t xt) {
	y := func(t *xt) {
		fmt.Println("in func")
		t.name = "Suxy_Q"
	}
	fmt.Println("struct x =", x)
	//y(x)
	y(&x)
	fmt.Println("struct x =", x)
}

func const_ex1() {
	const a = 3.14
	const b int32 = 186282
	fmt.Printf("%v %T %v %T\n", a, a, b, b)
	var c float64
	c = a * float64(b)
	fmt.Printf(" c= %v %T\n", c, c)

}

func runToy() {
	fmt.Println("in Toy")
	var t = toy.New("bear", 1)
	fmt.Printf("%v\n", t)
	t.Update(10, 5)
	fmt.Printf("%v\n", t)
	_, _, c, d := t.Get()
	fmt.Printf("The toy %s weighs %f, %d on hand and %d sold", t.Name, t.Weight, c, d)

}

func main() {

	var_ex1()
	struct_ex1()
	pointer_ex1()
	const_ex1()
	exercises.Methods()
	exercises.Embedding()
	exercises.Sport()
	runToy()

}
