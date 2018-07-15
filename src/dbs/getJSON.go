package dbs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

//
func GetTitle(w http.ResponseWriter, r *http.Request) {
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

//
func SetData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		var con Content
		fmt.Println(string(result))
		json.Unmarshal(result, &con)

		_, err := db.Exec("INSERT INTO content(category,name,results,number,searchcount) values(?,?,?,?,?)", con.Category, con.Name, con.Results, con.Number, 1)

		if err != nil {
			fmt.Fprintf(w, "{\"rtncode\":101}")
		} else {
			fmt.Fprintf(w, "{\"rtncode\":0}")
		}

	}
}

//
func QueryKey(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		var con Content
		var c ContentSlice
		c.RtnCode = 0
		fmt.Println(string(result))
		json.Unmarshal(result, &con)

		row := db.QueryRow("SELECT * FROM content where name = ? and results = ?", con.Name, con.Results)

		var tmp Content
		var ID int
		er := row.Scan(&ID, &tmp.Category, &tmp.Name, &tmp.Results, &tmp.Number, &tmp.SearchCount)
		if er != nil {
			if er == sql.ErrNoRows {
				//the select result is null
				c.RtnCode = 1
			} else {
				//select err
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
			//insert success,update the searchcount.
			updateCount(ID)
		}

	}
}
