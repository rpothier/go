package main

import "fmt"

func main() {

	defer fmt.Println("this is the last line")
	numSlice := [] int {5,4,3,2,1}

        fmt.Println(numSlice[2])

}



