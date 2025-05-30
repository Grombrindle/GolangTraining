package websocket

import (
	"fmt"
	http "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "helloworld ma nigga: %s\n", r.URL.Path)
	})
}
