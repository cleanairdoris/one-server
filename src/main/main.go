package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Title struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type TitleSlice struct {
	RtnCode int     `json:"rtncode"`
	Titles  []Title `json:"titles"`
}

type Content struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Number int64  `json:"number"`
}

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:6579.?a@tcp(localhost:3306)/one?charset=utf8")
}

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

func getData(w http.ResponseWriter, r *http.Request) {
	var t TitleSlice

	rows, err := db.Query("select * from title")
	if err != nil {
		t.RtnCode = 101
	}
	t.RtnCode = 1
	for rows.Next() {
		t.RtnCode = 0
		var s string
		var i int
		err = rows.Scan(&i, &s)
		if err != nil {
			t.RtnCode = 2
		} else {
			t.Titles = append(t.Titles, Title{ID: i, Name: s})
		}
	}

	rows.Close()
	b, err := json.Marshal(t)
	if err != nil {
		t.RtnCode = 3
	} else {
		fmt.Fprintf(w, "%s", b)
	}
}

func setData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		var con Content
		json.Unmarshal([]byte(result), &con)

		_, err := db.Exec("INSERT INTO content(type,name,value,number) values(?,?,?,?)", con.Type, con.Name, con.Value, 1)

		if err != nil {
			fmt.Fprintf(w, "%s", "insert fail")
		} else {
			fmt.Fprintf(w, "%s", "insert sucess")
		}

	}
}

func main() {
	http.HandleFunc("/getTitle", getTitle)
	http.HandleFunc("/getData", getData)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
