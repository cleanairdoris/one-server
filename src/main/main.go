package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getTitle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	titleFile := "title.json"
	buf, err := ioutil.ReadFile(titleFile)
	if err != nil {
		fmt.Fprintf(w, "{\"code\":101}")
	} else {
		fmt.Fprintf(w, "%s", buf)
	}

}

func main() {
	http.HandleFunc("/getTitle", getTitle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
