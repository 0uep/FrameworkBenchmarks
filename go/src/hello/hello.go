package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
)

type MessageStruct struct {
	Message string
}

type World struct {
	Id           uint16 `json:"id"`
	RandomNumber uint16 `json:"randomNumber"`
}

type Fortune struct {
	Id           uint16 `json:"id"`
	Message      uint16 `json:"message"`
}

const (
	DB_CONN_STR           = "benchmarkdbuser:benchmarkdbpass@tcp(172.16.98.98:3306)/hello_world?charset=utf8"
	DB_SELECT_SQL         = "SELECT id, randomNumber FROM World where id = ?;"
  DB_FORTUNE_SELECT_SQL = "SELECT id, message FROM Fortune;"
	DB_ROWS               = 10000
)

var (
	db    *sql.DB
	query *sql.Stmt
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var err error
	if db, err = sql.Open("mysql", DB_CONN_STR); err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	if query, err = db.Prepare(DB_SELECT_SQL); err != nil {
		log.Fatalf("Error preparing statement: %s", err)
	}
	if fortuneQuery, err = db.Query(DB_FORTUNE_SELECT_SQL); err != nil {
		log.Fatalf("Error preparing statement: %s", err)
	}
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/db", dbHandler)
  http.HandleFunc("/fortune", fortuneHandler)
	http.ListenAndServe(":8080", nil)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	j, _ := json.Marshal(&MessageStruct{"Hello, world"})
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.Write(j)
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
  	qnum := 1
	if qnumStr := r.URL.Query().Get("queries"); len(qnumStr) != 0 {
		qnum, _ = strconv.Atoi(qnumStr)
	}
	ww := make([]World, qnum)
	for i := 0; i < qnum; i++ {
		query.QueryRow(rand.Intn(DB_ROWS)+1).Scan(&ww[i].Id, &ww[i].RandomNumber)
	}
	w.Header().Set("Content-Type", "application/javascript")
	j, _ := json.Marshal(ww)
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.Write(j)
}

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
  fortunes := make([]Fortune, 13)
  
  // Execute the query
	rows, err := db.Query(DB_FORTUNE_SELECT_SQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %s", err)
	}
  
  var i := 0
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(&fortunes[i].id, &fortunes[i].message)
		if err != nil {
			panic(err.Error())
		}
		i++
	}
  
  var tmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/fortune.html"))
  if err := tmpl.Execute(w, map[string]interface{} {"fortunes": fortunes}); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
