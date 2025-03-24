package main

import (
	"go-flask-rewrite/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/info", handlers.GetInfo)
	http.HandleFunc("/reload_acl", handlers.ReloadACL)
	http.HandleFunc("/data_record", handlers.DataRecord)
	http.HandleFunc("/data_usage", handlers.DataUsage)

	http.ListenAndServe(":5000", nil)
}
