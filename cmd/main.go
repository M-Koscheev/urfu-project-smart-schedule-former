package main

import (
	_ "fmt"
	_ "net/http"
	_ "time"
)

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

func main() {
	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)
	test()
}
