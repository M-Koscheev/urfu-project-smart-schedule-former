package cmd

import (
	_ "fmt"
	_ "net/http"
	_ "time"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/prof_func"
)

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

func StartApp() {
	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)
	prof_func.Test()
}
