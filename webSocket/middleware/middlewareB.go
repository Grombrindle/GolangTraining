package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// so here i can make somthing like see if the token is correct and some error handling or maybe if i want i can do some logic or somthing
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func MiddlewareB() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	// http.ListenAndServe(":8080", nil)
}
