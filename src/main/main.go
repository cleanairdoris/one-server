package main

import (
	"dbs"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getTitle", dbs.GetTitle)
	http.HandleFunc("/setData", dbs.SetData)
	http.HandleFunc("/queryKey", dbs.QueryKey)

	err := http.ListenAndServe(":18080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
