package main

import "fmt"

import . "github.com/tj/go-debug"
import "time"

var a = Debug("goroutine:a")
var b = Debug("goroutine:b")
var c = Debug("goroutine:c")

func work(debug DebugFunction, delay time.Duration) {
	for {
		fmt.Println("test3")
		debug("doing stuff")
		time.Sleep(delay)
	}
}

func main() {

	defer fmt.Println("this is the last line")
	//q := make(chan bool)
	go work(a, 1000*time.Millisecond)
	go work(b, 250*time.Millisecond)
	go work(c, 100*time.Millisecond)
	fmt.Println("test4")
	time.Sleep(10*time.Second)
	//<-q

}
