package main

import (
	"Classical/Backend/controller"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

func main() {

	//db.Connect()
	classController := controller.NewClassController()
	postController := controller.NewPostController()

	router := mux.NewRouter()
	//class API endpoints and functionality
	router.HandleFunc("/getClasses", classController.GetClasses).Methods("GET")

	router.HandleFunc("/createClass", classController.CreateClass).Methods("POST")
	router.HandleFunc("/deleteClass/{className}", classController.DeleteClass).Methods("DELETE")
	router.HandleFunc("/getTrendingClasses", classController.GetSortedClasses).Methods("GET")
	router.HandleFunc("/getClassesByName/{className}", classController.GetClasessByName).Methods("GET")

	//post API endpoints and functionality
	router.HandleFunc("/createClassPost", postController.CreateClassPost).Methods("POST")
	router.HandleFunc("/getPostsByClassName/{className}", postController.GetClassPostsByName).Methods("GET")
	router.HandleFunc("/increasePostVotes", postController.IncreasePostVote).Methods("POST")
	router.HandleFunc("/decreasePostVotes", postController.DecreasePostVotes).Methods("POST")

	// Add CORS headers to all responses
	handler := cors.Default().Handler(router)

	http.ListenAndServe(":8000", handler)
}
