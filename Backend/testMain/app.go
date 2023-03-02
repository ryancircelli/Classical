package testMain

import (
	"Classical/Backend/controller"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	var err error

	a.DB, err = sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}

	err = PingDB(a)
	if err != nil {
		panic(err)
	}

	fmt.Println("CONNECTED")

	a.DB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	a.DB.SetMaxIdleConns(5)
	a.DB.SetMaxOpenConns(2)
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}
func PingDB(a *App) error {
	err := a.DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Run(addr string) {
	http.ListenAndServe(addr, a.Router)
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
