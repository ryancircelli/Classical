package controller

import (
	obj "Classical/Backend/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func GetClasses(w http.ResponseWriter, r *http.Request) {
	// var classes []obj.Class
	// result, err := db.DB.Query("SELECT * from class")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer result.Close()
	// for result.Next() {
	// 	var class obj.Class
	// 	err := result.Scan(&class.ID, &class.ClassName)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	classes = append(classes, class)
	// }
	// respondWithJSON(w, http.StatusOK, classes)

	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")

	if err != nil {
		panic(err)
	}
	//w.Header().Set("Content-Type", "application/json")
	var classes []obj.Class
	result, err := db.Query("SELECT * from class")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var class obj.Class
		err := result.Scan(&class.ID, &class.ClassName)
		if err != nil {
			panic(err.Error())
		}
		classes = append(classes, class)
	}
	//json.NewEncoder(w).Encode(classes)
	respondWithJSON(w, http.StatusOK, classes)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var class obj.Class
	err = decoder.Decode(&class)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO class(className) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var classes []obj.Class
	rows, err := db.Query("SELECT * FROM class WHERE className = ?", class.ClassName)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cla obj.Class
		if err := rows.Scan(&cla.ID, &cla.ClassName); err != nil {
			panic(err.Error())
		}
		classes = append(classes, cla)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	if len(classes) == 1 {
		fmt.Fprintf(w, "Class with Name = %s already exists", class.ClassName)
		return
	}

	res, err := stmt.Exec(class.ClassName)
	if err != nil {
		panic(err.Error())
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		id, _ := res.LastInsertId()
		class.ID = int64(id)
		respondWithJSON(w, http.StatusOK, class)
	}

}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	params := mux.Vars(r)
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
