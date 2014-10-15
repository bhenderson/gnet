package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhenderson/gnet"
)

func main() {
	l := gnet.Must(gnet.Listen("tcp", ":8080"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "hello world!")
	})

	log.Fatal(http.Serve(l, nil))
}
