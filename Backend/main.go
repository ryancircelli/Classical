package main

import (
	"Classical/Backend/controller"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
