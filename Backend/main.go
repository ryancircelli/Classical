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

	// router.HandleFunc("/posts", createPost).Methods("POST")
	// router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	// router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	http.ListenAndServe(":8000", router)

	// posts, err := postsByClassID(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("posts found: %v\n", posts)

}

// func postsByClassID(ID int) ([]Posts, error) {
// 	var posts []Posts

// 	rows, err := db.Query("SELECT * FROM posts WHERE classID = ?", ID)

// 	if err != nil {
// 		return nil, fmt.Errorf("postsByClassID %q: %v", ID, err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var pos Posts
// 		if err := rows.Scan(&pos.classID, &pos.postID, &pos.postName, &pos.postContent); err != nil {
// 			return nil, fmt.Errorf("PostsByClassID %q: %v", ID, err)
// 		}
// 		posts = append(posts, pos)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("postsByClassID %q: %v", ID, err)
// 	}
// 	return posts, nil
// }

// // addClass adds the specified class into the database and returns the classID of the class
// func addClassPost(post Posts) (int64, error) {
// 	result, err := db.Exec("INSERT INTO posts (classID, postName, postContent) VALUES (?,?,?)", post.classID, post.postName, post.postContent)

// 	if err != nil {
// 		return 0, fmt.Errorf("addClassPost: %v", err)
// 	}
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, fmt.Errorf("addClassPost: %v", err)
// 	}

// 	return id, nil
// }
