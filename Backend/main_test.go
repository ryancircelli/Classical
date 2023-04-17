// main_test.go
package main

import (
	"Classical/Backend/controller"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetClasses(t *testing.T) {
	req, err := http.NewRequest("GET", "/getClasses", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetClasses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"className":"COP5000"},{"id":2,"className":"CIS4930"},{"id":3,"className":"CGS3065"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateClass(t *testing.T) {
	var jsonStr = []byte(`{"className":"COP12341"}`)

	req, err := http.NewRequest("POST", "/createClass", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application.json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateClass)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected := `{"id":4,"className":"COP12341","total_votes":0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

}

// deleteClass test
func TestDeleteClass(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/deleteClass/COP3502", nil)
	if err != nil {
		panic(err.Error())
	}
	req = mux.SetURLVars(req, map[string]string{"className": "COP3502"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteClass)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Class with Name = COP3502 was deleted`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

// CreatePost test
func TestCreatePost(t *testing.T) {
	var jsonStr = []byte(`{"classID":1,"postName":"Discord Link 1","postContent":"www.DiscordLink1.com"}`)

	req, err := http.NewRequest("POST", "/createClassPost", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application.json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateClassPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected := `{"postId":4,"classId":1,"postName":"Discord Link 1","postContent":"www.DiscordLink1.com","postVotes":0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

}

// getPosts Test
func TestGetPostsByClassID(t *testing.T) {
	req, err := http.NewRequest("GET", "/getPostsByClassId/2", nil)
	if err != nil {
		panic(err.Error())
	}
	req = mux.SetURLVars(req, map[string]string{"classID": "2"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetClassPosts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"postId":3,"classId":2,"postName":"GroupMe Link","postContent":"www.groupme.com","postVotes":9}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

// IncreaseVotes test
// func TestIncreasePostVotes(t *testing.T) {
// 	req, err := http.NewRequest("PUT", "/increasePostVotes/4", nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	req = mux.SetURLVars(req, map[string]string{"postID": "4"})
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controller.IncreasePostVote)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `Post with ID = 1 was updated`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}

// }

// // DecreaseVotes test
// func TestDecreasePostVotes(t *testing.T) {
// 	req, err := http.NewRequest("PUT", "/decreasePostVotes/1", nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	req = mux.SetURLVars(req, map[string]string{"postID": "1"})
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controller.DecreasePostVotes)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `Post with ID = 1 was updated`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}

// }

func TestGetClassesByFirstLetters(t *testing.T) {
	req, err := http.NewRequest("GET", "/getClassesByFirstLetters/CEN3", nil)
	if err != nil {
		panic(err.Error())
	}
	req = mux.SetURLVars(req, map[string]string{"className": "CEN3"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetClasessByName)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":5,"className":"CEN3031"},{"id":6,"className":"CEN3032"},{"id":7,"className":"CEN3"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// getTrendingClasses test
func TestGetTrendingClass(t *testing.T) {
	req, err := http.NewRequest("GET", "/getTrendingClasses", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetSortedClasses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":2,"className":"CIS4930","total_votes":9},{"id":1,"className":"COP5000","total_votes":5},{"id":3,"className":"CGS3065","total_votes":0},{"id":4,"className":"COP12341","total_votes":0}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
