// main_test.go
package main

import (
	"Classical/Backend/testMain"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a testMain.App

func TestMain(m *testing.M) {
	a = testMain.App{}
	a.Initialize("root", "password123", "classical")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}
func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(tableCreationQuery2); err != nil {
		log.Fatal(err)
	}

}
func clearTable() {
	a.DB.Exec("DELETE FROM class")
	a.DB.Exec("DELETE FROM post")

	a.DB.Exec("ALTER TABLE class AUTO_INCREMENT = 1")
	a.DB.Exec("ALTER TABLE post AUTO_INCREMENT = 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/getClasses", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentPostsByClassId(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/getPostsByClassId/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Posts not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
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
