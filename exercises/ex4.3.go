/*
Given a list of names, you need to organize each name within a slice based on its length.
*/

package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {

	var size int = 0
	for n := range names {
		if len(names[n]) > size {
			size = len(names[n])
		}
	}
	//fmt.Println("largest name is", size)
	answer := make([][] string, size)

	//fmt.Println(names)

	for n := range names {
		//fmt.Println(names[n])
		//fmt.Println(len(names[n]))
		answer[len(names[n])-1] = append(answer[len(names[n])-1], names[n])
	}
	fmt.Println(answer)
}
