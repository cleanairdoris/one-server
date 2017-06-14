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
	Category    string `json:"category"`
	Name        string `json:"name"`
	Results     string `json:"results"`
	Number      int64  `json:"number"`
	SearchCount int64  `json:"searchcount"`
}

type ContentSlice struct {
	RtnCode  int       `json:"rtncode"`
	Contents []Content `json:"contents"`
}

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:6579.?a@tcp(localhost:3306)/one?charset=utf8")
}

func getTitle(w http.ResponseWriter, r *http.Request) {
	var t TitleSlice

	rows, err := db.Query("select * from title")
	if err != nil {
		t.RtnCode = 101
	}
	t.RtnCode = 1
	for rows.Next() {
		t.RtnCode = 0

		var tmp Title
		err = rows.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			t.RtnCode = 2
		} else {
			t.Titles = append(t.Titles, tmp)
		}
	}

	rows.Close()
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Fprintf(w, "{\"rtncode\":1101}")
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
		fmt.Println(string(result))
		json.Unmarshal(result, &con)

		_, err := db.Exec("INSERT INTO content(type,name,value,number) values(?,?,?,?)", con.Category, con.Name, con.Results, con.Number)

		if err != nil {
			fmt.Fprintf(w, "{\"rtncode\":0}")
		} else {
			fmt.Fprintf(w, "{\"rtncode\":101}")
		}

	}
}

func queryKey(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		var con Content
		var c ContentSlice
		c.RtnCode = 0
		fmt.Println(string(result))
		json.Unmarshal(result, &con)

		row := db.QueryRow("SELECT * FROM content where name = ? and value = ?", con.Name, con.Results)

		var tmp Content
		var ID int
		er := row.Scan(&ID, &tmp.Category, &tmp.Name, &tmp.Results, &tmp.Number, &tmp.SearchCount)
		if er != nil {
			if er == sql.ErrNoRows {
				c.RtnCode = 1
			} else {
				c.RtnCode = 2
			}

		} else {
			c.RtnCode = 0
			c.Contents = append(c.Contents, tmp)
		}

		b, err := json.Marshal(c)
		if err != nil {
			fmt.Fprintf(w, "{\"rtncode\":1101}")
		} else {
			fmt.Fprintf(w, "%s", b)
			updateCount(ID)
		}

	}
}

func updateCount(ID int) {
	db.Exec("update content set searchcount = searchcount + 1 where id = ?", ID)
}

func main() {
	http.HandleFunc("/getTitle", getTitle)
	http.HandleFunc("/setData", setData)
	http.HandleFunc("/queryKey", queryKey)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
