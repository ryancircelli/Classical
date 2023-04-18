package controller

import (
	obj "Classical/Backend/model"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func GetClasses(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var classes []obj.Class
	result, err := db.Query("SELECT className, lastUpdated from class")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var class obj.Class
		err := result.Scan(&class.ClassName, &class.LastUpdated)
		if err != nil {
			panic(err.Error())
		}
		classes = append(classes, class)
	}

	// Convert the classes to ClassWithoutTotalVotes
	classesWithoutTotalVotes := make([]obj.ClassWithoutTotalVotes, len(classes))
	for i, class := range classes {
		classesWithoutTotalVotes[i] = obj.ClassWithoutTotalVotes{Class: class}
	}

	respondWithJSON(w, http.StatusOK, classesWithoutTotalVotes)
}

func GetClasessByName(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
	if err != nil {
		panic(err)
	}

	params := mux.Vars(r)
	searchTerm := params["className"]
	var result *sql.Rows

	// Check if the search term is a class number
	match, err := regexp.MatchString(`^\d+$`, searchTerm)
	if err != nil {
		panic(err.Error())
	}
	if match {
		// Search by class number
		result, err = db.Query("SELECT className, lastUpdated, totalVotes FROM class WHERE className REGEXP ? ORDER BY LENGTH(className), className", searchTerm+"[[:digit:]]*")
	} else {
		// Search by class name
		result, err = db.Query("SELECT className, lastUpdated, totalVotes from class WHERE className LIKE ? ORDER BY LENGTH(className), className", searchTerm+"%")
	}

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var classes []obj.Class

	for result.Next() {
		var class obj.Class
		err := result.Scan(&class.ClassName, &class.LastUpdated, &class.TotalVotes)
		if err != nil {
			panic(err.Error())
		}
		classes = append(classes, class)
	}

	// Check if there are any classes
	if len(classes) == 0 {
		respondWithJSON(w, http.StatusNotFound, []obj.ClassWithoutTotalVotes{})
	} else {
		jsonData, err := json.Marshal(classes)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

func GetSortedClasses(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `
		SELECT c.className, c.lastUpdated, COALESCE(SUM(p.postVotes), 0) as total_votes
		FROM class c
		LEFT JOIN post p ON c.className = p.postClassName
		GROUP BY c.className
		ORDER BY total_votes DESC;
	`
	//Figure this out
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	classes := make([]obj.Class, 0)

	for rows.Next() {
		var class obj.Class
		err := rows.Scan(&class.ClassName, &class.LastUpdated, &class.TotalVotes)
		if err != nil {
			log.Fatal(err)
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, http.StatusOK, classes)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var class obj.Class
	err = decoder.Decode(&class)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	stmt, err := db.Prepare("INSERT INTO class(className) VALUES(?)")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer stmt.Close()

	var classes []obj.Class
	rows, err := db.Query("SELECT * FROM class WHERE className = ?", class.ClassName)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cla obj.Class
		if err := rows.Scan(&cla.ClassName); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		classes = append(classes, cla)
	}
	if err := rows.Err(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(classes) == 1 {
		err := errors.New("Class with Name = " + class.ClassName + " already exists")
		respondWithError(w, http.StatusConflict, err.Error())
		return
	}

	res, err := stmt.Exec(class.ClassName)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		respondWithJSON(w, http.StatusOK, class)
	}
}

func DeleteClass(w http.ResponseWriter, r *http.Request) { //Figure this out
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	params := mux.Vars(r)

	// Delete all associated posts before deleting the class
	_, err = db.Exec("DELETE FROM post WHERE postClassName = (SELECT className FROM class WHERE className = ?)", params["className"])
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("DELETE FROM class WHERE className = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["className"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Class with Name = %s was deleted", params["className"])
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//encode payload to json
	response, _ := json.Marshal(payload)
	// set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]interface{}{"error": message}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode error message", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
