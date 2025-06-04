package main

import (
	"fmt"
	sql "helloworld/Sql"
	gg "helloworld/modules"

	// web "helloworld/webSocket"
	"log"
	// "net/http"
	// "rsc.io/quote"
)

func Go() string {
	return "kaka wasaka"
}

func main() {
	sql.DatabaseTest()

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.

	// web.HttpCall()

	for i := 0; i < 10; i++ {
		println("")
	}

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
