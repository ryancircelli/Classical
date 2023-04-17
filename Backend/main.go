package main

import (
	"Classical/Backend/controller"
	obj "Classical/Backend/model"
	"database/sql"
	"encoding/json"
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
	router.HandleFunc("/getPostsByClassId/{classID}", controller.GetClassPosts).Methods("GET")

	router.HandleFunc("/getPostsByClassId/{classID}", controller.GetClassPosts).Methods("GET")

	//API call for post votes
	router.HandleFunc("/increasePostVotes/{postID}", controller.IncreasePostVote).Methods("PUT")
	router.HandleFunc("/decreasePostVotes/{postID}", controller.DecreasePostVotes).Methods("PUT")
	router.HandleFunc("/getClassesByName/{className}", controller.GetClassByName).Methods("GET")
	router.HandleFunc("/getClassesByFirstLetters/{className}", controller.GetClassByFirstThreeLetters).Methods("GET")
	router.HandleFunc("/getTrendingClasses", controller.GetSortedClasses).Methods("GET")

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

func GetClassesTest(w http.ResponseWriter, r *http.Request) {
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

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//encode payload to json
	response, _ := json.Marshal(payload)

	// set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
