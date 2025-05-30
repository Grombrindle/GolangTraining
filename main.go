package main

import (
	"fmt"
<<<<<<< HEAD
	web "helloworld/Websocket"
	gg "helloworld/modules"
	"log"
	// "net/http"
=======
	"log"

	gg "helloworld/modules"
>>>>>>> 291e817528becc56ad1799d8ac451a3f7bf3aa20
	// "rsc.io/quote"
)

func Go() string {
	return "kaka wasaka"
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
<<<<<<< HEAD

	web.HttpCall()

	for i := 0; i < 10; i++ {
		println("")
	}

=======
>>>>>>> 291e817528becc56ad1799d8ac451a3f7bf3aa20
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var seq []uint64
	fmt.Println(gg.FibSeqRecursive(7, seq))

	// Request a greeting message.
	message, err := gg.Hello("dasdsad")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
