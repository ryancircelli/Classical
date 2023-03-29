package controller

import (
	obj "Classical/Backend/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateClassPost(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var post obj.Post
	err = decoder.Decode(&post)
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare("INSERT INTO post(classId, postName, postContent) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(post.ClassID, post.PostName, post.PostContent)
	if err != nil {
		panic(err.Error())
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		id, _ := res.LastInsertId()
		post.PostID = int64(id)
		respondWithJSON(w, http.StatusOK, post)
	}

	// stmt, err := db.Prepare("INSERT INTO post(classId, postName, postContent) VALUES(?, ?, ?)")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// keyVal := make(map[string]string)
	// json.Unmarshal(body, &keyVal)
	// classId := keyVal["classID"]
	// postName := keyVal["postName"]
	// postContent := keyVal["postContent"]
	// intVar, err := strconv.Atoi(classId)
	// _, err = stmt.Exec(intVar, postName, postContent)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// f.Fprintf(w, "New post was created")
}

func IncreasePostVote(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE post SET postVotes = postVotes + 1 WHERE postID = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	_, err = stmt.Exec(params["postID"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params["postID"])
}

func DecreasePostVotes(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE post SET postVotes = postVotes - 1 WHERE postID = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	_, err = stmt.Exec(params["postID"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params["postID"])
}

func GetClassPosts(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var posts []obj.Post
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM post WHERE classID = ?", params["classID"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post obj.Post
	for result.Next() {
		err := result.Scan(&post.PostID, &post.ClassID, &post.PostName, &post.PostContent, &post.PostVotes)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	if len(posts) == 0 {
		fmt.Fprintf(w, "Class with id = %v does not have any posts", post.ClassID)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
