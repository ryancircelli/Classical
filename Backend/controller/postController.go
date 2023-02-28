package controller

import (
	"Classical/Backend/db"
	obj "Classical/Backend/model"
	"encoding/json"
	f "fmt"
	"io/ioutil"
	"net/http"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateClassPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.DB.Prepare("INSERT INTO post(classId, postName, postContent) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	classId := keyVal["classID"]
	postName := keyVal["postName"]
	postContent := keyVal["postContent"]
	intVar, err := strconv.Atoi(classId)
	_, err = stmt.Exec(intVar, postName, postContent)
	if err != nil {
		panic(err.Error())
	}
	f.Fprintf(w, "New post was created")
}

func GetClassPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []obj.Post
	params := mux.Vars(r)
	result, err := db.DB.Query("SELECT * FROM post WHERE classID = ?", params["classID"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post obj.Post
	for result.Next() {
		err := result.Scan(&post.PostID, &post.ClassID, &post.PostName, &post.PostContent)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
