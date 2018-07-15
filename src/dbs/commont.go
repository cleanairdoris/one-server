package dbs

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	//init mysql
	_ "github.com/go-sql-driver/mysql"
)

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
	var config configs
	json.Unmarshal(buf, &config)

	//db, _ = sql.Open("mysql", "root:6579.?a@tcp(localhost:3306)/one?charset=utf8")
	db, _ = sql.Open(config.DbType, config.User+":"+config.Pwd+"@tcp(localhost:3306)/"+config.DbName+"?charset=utf8")
}

func updateCount(ID int) {
	db.Exec("update content set searchcount = searchcount + 1 where id = ?", ID)
}
