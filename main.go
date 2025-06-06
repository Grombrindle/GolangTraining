package main

import (
	"fmt"
	sql "helloworld/Sql"
	gg "helloworld/modules"
	tmpt "helloworld/webSocket"
	middleware "helloworld/webSocket/middleware"

	// web "helloworld/webSocket"
	"log"
	// "net/http"
	// "rsc.io/quote"
)

func Go() string {
	return "kaka wasaka"
}

func main() {
	// Start templates in a goroutine so it doesn't block
	go tmpt.WebSocketTest()
	go tmpt.JsonUser()
	go tmpt.SessionCookie()
	go tmpt.Templates()
	go tmpt.Form()
	go middleware.Hash()
	go middleware.MiddlewareB()
	go middleware.MiddlewareA()
	// Start database test in a goroutine
	go sql.DatabaseTest()

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
	// allow the program to continue.
	if err != nil {
		log.Println("Error calling gg.Hello:", err)
	} else {
		// If no error was returned, print the returned message
		// to the console.
		fmt.Println(message)
	}

	// Keep main thread running to allow goroutines to continue
	select {}
}
