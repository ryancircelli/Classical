package main

import (
	"Classical/Backend/controller"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
	// Set the allowed methods in the CORS options
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	handler := corsMiddleware(router)
	http.ListenAndServe(":8000", handler)
}
