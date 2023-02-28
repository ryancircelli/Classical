package controller

import (
	"Classical/Backend/db"
	obj "Classical/Backend/model"
	"Classical/Backend/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func GetClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var classes []obj.Class
	result, err := db.DB.Query("SELECT * from class")
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
	json.NewEncoder(w).Encode(classes)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.DB.Prepare("INSERT INTO class(className) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	className := keyVal["className"]

	//Make sure class does not already exist
	classesCheckArray, err := service.ClassesByName(className)
	if len(classesCheckArray) == 1 {
		fmt.Fprintf(w, "Class with Name = %s already exists", className)
		return
	}
	_, err = stmt.Exec(className)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New class was created")
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.DB.Prepare("DELETE FROM class WHERE className = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["className"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Class with Name = %s was deleted", params["className"])
}
