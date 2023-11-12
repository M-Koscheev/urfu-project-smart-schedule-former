package cmd

import (
	"fmt"
	"net/http"

	//"time"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/database_func"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/prof_func"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, database_func.GetKnowledgeList())
}

func StartApp() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
	prof_func.Test()
}
