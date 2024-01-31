package main

import (
	"fmt"
	"log"
	"example.com/greeting"
	"rsc.io/quote"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greeting.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
