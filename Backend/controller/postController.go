package controller

import (
	obj "Classical/Backend/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type PostController interface {
	CreateClassPost(http.ResponseWriter, *http.Request)
	IncreasePostVote(w http.ResponseWriter, r *http.Request)
	DecreasePostVotes(w http.ResponseWriter, r *http.Request)
	GetClassPostsByName(w http.ResponseWriter, r *http.Request)
}

type postControllerImpl struct{}

func NewPostController() PostController {
	return &postControllerImpl{}
}

func (c *postControllerImpl) CreateClassPost(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
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
	stmt, err := db.Prepare("INSERT INTO post(postClassName, postName, postContent) VALUES(?, ?, ?)")
	stmt2, err2 := db.Prepare("UPDATE class SET lastUpdated = CURRENT_TIMESTAMP WHERE className = ?")
	if err != nil {
		panic(err.Error())
	}
	if err2 != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(post.PostClassName, post.PostName, post.PostContent)
	_, err2 = stmt2.Exec(post.PostClassName)
	if err != nil {
		panic(err.Error())
	}
	if err2 != nil {
		panic(err.Error())
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		id, _ := res.LastInsertId()
		post.PostID = int64(id)
		var returnPost obj.Post
		result, err := db.Query("SELECT * FROM post WHERE postClassName = ?", post.PostClassName)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()
		for result.Next() {
			err := result.Scan(&returnPost.PostID, &returnPost.PostClassName, &returnPost.PostName, &returnPost.PostContent, &returnPost.PostVotes, &returnPost.TimePosted)
			if err != nil {
				panic(err.Error())
			}
		}
		respondWithJSON(w, http.StatusOK, returnPost)
	}
}

func (c *postControllerImpl) IncreasePostVote(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
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
	stmt, err := db.Prepare("UPDATE post SET postVotes = postVotes + 1 WHERE postID = ?")
	stmt2, err2 := db.Prepare("UPDATE class SET lastUpdated = CURRENT_TIMESTAMP, totalVotes = totalVotes +1 WHERE className = ?")
	if err != nil {
		panic(err.Error())
	}
	if err2 != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(post.PostID)
	_, err2 = stmt2.Exec(post.PostClassName)
	if err != nil {
		panic(err.Error())
	}
	if err2 != nil {
		panic(err.Error())
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		fmt.Fprintf(w, "Post with id = %v was updated", post.PostID)
	}

}

func (c *postControllerImpl) DecreasePostVotes(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
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
	stmt, err := db.Prepare("UPDATE post SET postVotes = postVotes - 1 WHERE postID = ?")
	stmt2, err2 := db.Prepare("UPDATE class SET lastUpdated = CURRENT_TIMESTAMP, totalVotes = totalVotes - 1 WHERE className = ?")
	if err != nil {
		panic(err.Error())
	}
	if err2 != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(post.PostID)
	_, err2 = stmt2.Exec(post.PostClassName)
	if err != nil {
		panic(err.Error())
	}
	if err2 != nil {
		panic(err.Error())
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		fmt.Fprintf(w, "Post with id = %v was updated", post.PostID)
	}

}

func (c *postControllerImpl) GetClassPostsByName(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var posts []obj.Post
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM post WHERE postClassName = ?", params["className"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post obj.Post
	for result.Next() {
		err := result.Scan(&post.PostID, &post.PostClassName, &post.PostName, &post.PostContent, &post.PostVotes, &post.TimePosted)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	if len(posts) == 0 {
		fmt.Fprintf(w, "Class with name = %v does not have any posts", post.PostClassName)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
