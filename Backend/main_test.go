// main_test.go
package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var DB *sql.DB
var router *mux.Router

func TestMain(m *testing.M) {

	DB, err = sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")

	if err != nil {
		log.Fatal(err)
	}

	err := DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")

	router = mux.NewRouter()

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := DB.Exec(tableCreationQuery2); err != nil {
		log.Fatal(err)
	}

}

func clearTable() {
	DB.Exec("DELETE FROM class")
	DB.Exec("DELETE FROM post")

	DB.Exec("ALTER TABLE class AUTO_INCREMENT = 1")
	DB.Exec("ALTER TABLE post AUTO_INCREMENT = 1")
}

func TestEmptyTable(t *testing.T) {
	req, _ := http.NewRequest("GET", "localhost:8000/getClasses", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestAddClass(t *testing.T) {
	payload := []byte(`{"className":"COP5000}`)
	req, _ := http.NewRequest("POST", "localhost:8000/createClass", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	if m["className"] != "COP5000" {
		t.Errorf("expected name to be COP5000 got %v", m["className"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS class (
  id         INT AUTO_INCREMENT NOT NULL,
  className  VARCHAR(128) NOT NULL,
  PRIMARY KEY (` + "`" + `id` + "`" + `)
);`

const tableCreationQuery2 = `
CREATE TABLE IF NOT EXISTS post (
	postID    INT AUTO_INCREMENT NOT NULL,
	classID   INT,
	FOREIGN KEY (classID) REFERENCES class(id),
	postName  VARCHAR(128) NOT NULL,
	postContent VARCHAR(128) NOT NULL,
	PRIMARY KEY (` + "`" + `postID` + "`" + `)
  );`
