package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("http", ":8888", "HTTP port number.")

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		value := r.FormValue("key")
		fmt.Println(value)
		w.Write([]byte(value))
	})
}

func main() {
	fmt.Println("Start HTTP Server localhost:", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
