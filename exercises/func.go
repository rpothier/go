package main

import (
	"fmt"
)

type mystruct struct {
	name  string
	value int
}

func getValue() int {
	return 7
}

func main() {

	a := mystruct{name: "rob",
		value: getValue(),
	}

	//	value: func () int {
	//	    return 8
	//	}
	//	}

	fmt.Println("Hello, playground", a)
}
