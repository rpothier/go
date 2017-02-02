package main

import (
	"fmt"
	"log"
	"net/http"
	//"time"

)

func main() {

	var input string
	fmt.Print("Type any key to exit: ")

	fmt.Print("Start Go routine: ")
	//go log.Fatal(http.ListenAndServe(":8080", router))
	//go http.ListenAndServe(":8080", router)


	go listen_and_serve()

	fmt.Scanln(&input)
	//time.Sleep(1 * time.Minute)
	fmt.Print("Exiting ...")

}


func listen_and_serve() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
