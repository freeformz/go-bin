// go build -o ../bin/web

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
