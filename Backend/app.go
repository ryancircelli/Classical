package main

import (
	"Classical/Backend/controller"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	//class API endpoints and functionality
	a.Router.HandleFunc("/getClasses", controller.GetClasses).Methods("GET")
	a.Router.HandleFunc("/createClass", controller.CreateClass).Methods("POST")
	a.Router.HandleFunc("/deleteClass/{className}", controller.DeleteClass).Methods("DELETE")
	//post API endpoints and functionality
	a.Router.HandleFunc("/createClassPost", controller.CreateClassPost).Methods("POST")
	a.Router.HandleFunc("/getPostsByClassId/{classID}", controller.GetClassPosts).Methods("GET")
}
