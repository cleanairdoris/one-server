package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"flag"

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

type Configs struct {
	DbType string `json:"dbtype"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	DbName string `json:"dbname"`
}

var db = &sql.DB{}

func init() {
	configFile := flag.String("cfg", "config.json", "the db config filename")
	fi, err := os.Open(*configFile)
	fmt.Println(*configFile)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf, _ := ioutil.ReadAll(fi)
	var config Configs
	json.Unmarshal(buf, &config)

	//db, _ = sql.Open("mysql", "root:6579.?a@tcp(localhost:3306)/one?charset=utf8")
	db, _ = sql.Open(config.DbType, config.User+":"+config.Pwd+"@tcp(localhost:3306)/"+config.DbName+"?charset=utf8")
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

		_, err := db.Exec("INSERT INTO content(category,name,results,number,searchcount) values(?,?,?,?)", con.Category, con.Name, con.Results, con.Number, 1)

		if err != nil {
			fmt.Fprintf(w, "{\"rtncode\":101}")
		} else {
			fmt.Fprintf(w, "{\"rtncode\":0}")
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

func updateCount(ID int) {
	db.Exec("update content set searchcount = searchcount + 1 where id = ?", ID)
}

func prototest(w http.ResponseWriter, r *http.Request) {
	data := TestPro()
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Fprintf(w, "%s", str)
}

func main() {
	http.HandleFunc("/getTitle", getTitle)
	http.HandleFunc("/setData", setData)
	http.HandleFunc("/queryKey", queryKey)
	http.HandleFunc("/getproto", prototest)
	err := http.ListenAndServe(":18080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
