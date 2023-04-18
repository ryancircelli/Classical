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

	router := mux.NewRouter()
	//class API endpoints and functionality
	router.HandleFunc("/getClasses", controller.GetClasses).Methods("GET")
	router.HandleFunc("/createClass", controller.CreateClass).Methods("POST")
	router.HandleFunc("/deleteClass/{className}", controller.DeleteClass).Methods("DELETE")
	//post API endpoints and functionality
	router.HandleFunc("/createClassPost", controller.CreateClassPost).Methods("POST")
	router.HandleFunc("/getPostsByClassName/{className}", controller.GetClassPostsByName).Methods("GET")

	//API call for post votes
	router.HandleFunc("/increasePostVotes", controller.IncreasePostVote).Methods("POST")
	router.HandleFunc("/decreasePostVotes", controller.DecreasePostVotes).Methods("POST")
	router.HandleFunc("/getClassesByName/{className}", controller.GetClasessByName).Methods("GET")
	router.HandleFunc("/getTrendingClasses", controller.GetSortedClasses).Methods("GET")

	// router.HandleFunc("/posts", createPost).Methods("POST")
	// router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	// router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")

	// Add CORS headers to all responses
	handler := cors.Default().Handler(router)

	http.ListenAndServe(":8000", handler)

	// posts, err := postsByClassID(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("posts found: %v\n", posts)

}
