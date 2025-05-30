<<<<<<< HEAD
package Websocket

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	http "net/http"
	"time"

	"bytes"

	"github.com/gorilla/mux"
	// route "gorilla/mux"
)

func Post() {
	time.Sleep(1 * time.Second)
	title := "some_title"
	url := fmt.Sprintf("http://localhost:8000/books/%s", title)
	bodyData := map[string]string{"title": title}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	// vars := mux.Vars()
	// title := "da"
	// resp =  title
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status, title)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println(string(body))
}
func HttpCall() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("/static"))
	// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "helloworld ma nigga: %s\n", r.URL.Path)
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.Handle("/static", http.StripPrefix("/static", fs))
	go http.ListenAndServe(":8000", r)
	Post()
	// fmt.Println("dsadsa")
	// Restrict the request handler to specific hostnames or subdomains.

	// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	// Restrict the request handler to http/https.

	// r.HandleFunc("/secure", SecureHandler).Schemes("https")
	// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// Restrict the request handler to specific path prefixes.

	// 	bookrouter := r.PathPrefix("/books").Subrouter()
	// bookrouter.HandleFunc("/", AllBooks)
	// bookrouter.HandleFunc("/{title}", GetBook)

	// fs := http.FileServer(http.Dir("/static"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "helloworld ma nigga: %s\n", r.URL.Path)
	// 	// fmt.Fprintf(w, "shit is: %s\n", r.URL.Query().Get("token"))
	// 	// fmt.Print(r.URL.Query().Get("token"))
	// })
=======
package websocket

import (
	"fmt"
	http "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "helloworld ma nigga: %s\n", r.URL.Path)
	})
>>>>>>> 291e817528becc56ad1799d8ac451a3f7bf3aa20
}
