package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q %s", html.EscapeString(r.URL.Path), os.Getenv("DBHOST"))
	})

	log.Fatal(http.ListenAndServe(":9000", nil))

}

func Sum(x int, y int) int {
	return x + y
}
