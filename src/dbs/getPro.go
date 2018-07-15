package dbs

import (
	"database/sql"
	"dataprotobuf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func getTitlePro(w http.ResponseWriter, r *http.Request) {

	tts := &dataprotobuf.Titlelist{}
	rows, err := db.Query("select * from title")
	if err != nil {
		tts.Rtncode = 101
	}
	tts.Rtncode = 1
	for rows.Next() {
		tts.Rtncode = 0

		tmp := &dataprotobuf.Title{}
		err = rows.Scan(&tmp.Id, &tmp.Name)
		if err != nil {
			tts.Rtncode = 2
		} else {
			tts.Titles = append(tts.Titles, tmp)
		}
	}

	rows.Close()

	data, err := proto.Marshal(tts)
	if err != nil {
		fmt.Fprintf(w, "{\"rtncode\":1101}")
	} else {
		fmt.Fprintf(w, "%s", data)
	}
}

func setDataPro(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		con := &dataprotobuf.Content{}
		proto.Unmarshal(result, con)

		_, err := db.Exec("INSERT INTO pre_content(category,name,results,number,searchcount) values(?,?,?,?,?)", con.Category, con.Name, con.Results, con.Number, 1)

		if err != nil {
			fmt.Fprintf(w, "{\"rtncode\":101}")
		} else {
			fmt.Fprintf(w, "{\"rtncode\":0}")
		}

	}
}

func queryKeyPro(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		con := &dataprotobuf.Content{}
		c := &dataprotobuf.Contentlist{}
		c.Rtncode = 0
		fmt.Println(string(result))
		json.Unmarshal(result, &con)

		row := db.QueryRow("SELECT * FROM content where name = ? and results = ?", con.Name, con.Results)

		tmp := &dataprotobuf.Content{}
		var ID int
		er := row.Scan(&ID, &tmp.Category, &tmp.Name, &tmp.Results, &tmp.Number, &tmp.Searchcount)
		if er != nil {
			if er == sql.ErrNoRows {
				//the select result is null
				c.Rtncode = 1
			} else {
				//select err
				c.Rtncode = 2
			}

		} else {
			c.Rtncode = 0
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
