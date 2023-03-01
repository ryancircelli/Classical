package main

import (
	"Classical/Backend/controller"
	"Classical/Backend/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var err error

func main() {

	db.Connect()

	router := mux.NewRouter()
	//class API endpoints and functionality
	router.HandleFunc("/getClasses", controller.GetClasses).Methods("GET")
	router.HandleFunc("/createClass", controller.CreateClass).Methods("POST")
	router.HandleFunc("/deleteClass/{className}", controller.DeleteClass).Methods("DELETE")
	//post API endpoints and functionality
	router.HandleFunc("/createClassPost", controller.CreateClassPost).Methods("POST")
	router.HandleFunc("/getPostsByClassId/{classID}", controller.GetClassPosts).Methods("GET")

	// router.HandleFunc("/posts", createPost).Methods("POST")
	// router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	// router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")

	// Add CORS headers to all responses
	headers := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			next.ServeHTTP(w, r)
		})
	}
	router.Use(headers)

	http.ListenAndServe(":8000", router)

	// posts, err := postsByClassID(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("posts found: %v\n", posts)

}
