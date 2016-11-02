package main

import (
	"fmt"
	"strings"
	//"code.google.com/p/go-tour/wc"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	fmt.Println(s)
	a := strings.Split(s, " ")

	m := make(map[string]int)

	for i := range a {
		//fmt.Println(i)
		//fmt.Printf("word %d is %s\n", i, a[i])
		//elem, ok := m[a[i]]
		elem := m[a[i]]
		//fmt.Printf("type %T\n",elem)
		//fmt.Println(elem)
		//fmt.Println(ok)
		m[a[i]] = elem + 1
	}

	//fmt.Println(m)
	return m
}

func main() {

	WordCount("this is a test this")
	wc.Test(WordCount)
}
