package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		value := r.FormValue("key")
		fmt.Println(value)
		w.Write([]byte(value))
	})

	fmt.Println(http.ListenAndServe(":8888", nil))
}
