package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/intelfike/checkmodfile"
)

var port = flag.String("http", ":8888", "HTTP port number.")

func init() {
	f, err := checkmodfile.RegistFile("main.go")
	if err != nil {
		fmt.Print(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := f.GetLatest()
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		w.Write(b)
	})
}

func main() {
	fmt.Printf("Start HTTP Server localhost%s", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
