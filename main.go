package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("http", ":8888", "HTTP port number.")

func init() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ここにコード
	})
}

func main() {
	fmt.Printf("Start HTTP Server localhost:%s", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
