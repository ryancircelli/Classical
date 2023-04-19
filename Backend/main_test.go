// main_test.go
package main

import (
	"Classical/Backend/controller"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/gorilla/mux"
)

type Class struct {
	ID         int    `json:"id"`
	ClassName  string `json:"className"`
	TotalVotes int    `json:"total_votes" default:"0"`
}

type Post struct {
	PostID      int64  `json:"postId"`
	PostClassID int    `json:"classId"`
	PostName    string `json:"postName"`
	PostContent string `json:"postContent"`
	PostVotes   int    `json:"postVotes"`
}

func TestGetClasses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := controller.NewMockClassController(ctrl)

	mockClasses := []Class{
		{ID: 1, ClassName: "COP5000"},
		{ID: 2, ClassName: "CIS4930"},
		{ID: 3, ClassName: "CGS3065"},
	}

	mockController.
		EXPECT().
		GetClasses(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockClasses)
		})

	req, err := http.NewRequest("GET", "/getClasses", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockController.GetClasses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response into a slice of Class
	var returnedClasses []Class
	err = json.Unmarshal(rr.Body.Bytes(), &returnedClasses)
	if err != nil {
		t.Fatal(err)
	}

	// Define the expected data
	expectedClasses := []Class{
		{ID: 1, ClassName: "COP5000"},
		{ID: 2, ClassName: "CIS4930"},
		{ID: 3, ClassName: "CGS3065"},
	}

	// Compare the data
	if !reflect.DeepEqual(returnedClasses, expectedClasses) {
		t.Errorf("handler returned unexpected data: got %v want %v", returnedClasses, expectedClasses)
	}
}

func TestCreateClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := controller.NewMockClassController(ctrl)

	newClass := Class{ID: 4, ClassName: "COP12341"}
	jsonStr, err := json.Marshal(newClass)
	if err != nil {
		t.Fatal(err)
	}

	mockController.
		EXPECT().
		CreateClass(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonStr)
		})

	req, err := http.NewRequest("POST", "/createClass", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockController.CreateClass)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := string(jsonStr)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}

// deleteClass test
func TestDeleteClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := controller.NewMockClassController(ctrl)

	className := "COP3502"

	mockController.
		EXPECT().
		DeleteClass(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Class with Name = %s was deleted", className)))
		})

	req, err := http.NewRequest("DELETE", "/deleteClass/"+className, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"className": className})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockController.DeleteClass)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := fmt.Sprintf("Class with Name = %s was deleted", className)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetTrendingClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := controller.NewMockClassController(ctrl)

	mockClasses := []Class{
		{ID: 2, ClassName: "CIS4930", TotalVotes: 9},
		{ID: 1, ClassName: "COP5000", TotalVotes: 5},
		{ID: 3, ClassName: "CGS3065", TotalVotes: 0},
		{ID: 4, ClassName: "COP12341", TotalVotes: 0},
	}

	mockController.
		EXPECT().
		GetSortedClasses(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockClasses)
		})

	req, err := http.NewRequest("GET", "/getTrendingClasses", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockController.GetSortedClasses)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response into a slice of Class
	var returnedClasses []Class
	err = json.Unmarshal(rr.Body.Bytes(), &returnedClasses)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the data
	if !reflect.DeepEqual(returnedClasses, mockClasses) {
		t.Errorf("handler returned unexpected data: got %v want %v", returnedClasses, mockClasses)
	}
}

func TestGetClasessByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := controller.NewMockClassController(ctrl)

	mockClasses := []Class{
		{ID: 1, ClassName: "COP5000"},
		{ID: 2, ClassName: "CIS4930"},
	}

	mockController.
		EXPECT().
		GetClasessByName(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockClasses)
		})

	req, err := http.NewRequest("GET", "/getClassesByName/COP", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockController.GetClasessByName)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response into a slice of Class
	var returnedClasses []Class
	err = json.Unmarshal(rr.Body.Bytes(), &returnedClasses)
	if err != nil {
		t.Fatal(err)
	}

	// Define the expected data
	expectedClasses := []Class{
		{ID: 1, ClassName: "COP5000"},
		{ID: 2, ClassName: "CIS4930"},
	}

	// Compare the data
	if !reflect.DeepEqual(returnedClasses, expectedClasses) {
		t.Errorf("handler returned unexpected data: got %v want %v", returnedClasses, expectedClasses)
	}
}

// CreatePost test
func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := controller.NewMockPostController(ctrl)

	mockPost := Post{
		PostID:      4,
		PostClassID: 1,
		PostName:    "Discord Link 1",
		PostContent: "www.DiscordLink1.com",
		PostVotes:   0,
	}

	mockController.
		EXPECT().
		CreateClassPost(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockPost)
		})

	jsonStr := []byte(`{"classId":1,"postName":"Discord Link 1","postContent":"www.DiscordLink1.com"}`)
	req, err := http.NewRequest("POST", "/createClassPost", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockController.CreateClassPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `{"postId":4,"classId":1,"postName":"Discord Link 1","postContent":"www.DiscordLink1.com","postVotes":0}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}

// getPosts Test
func TestGetPostsByClassName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostController := controller.NewMockPostController(ctrl)

	className := "COP4600"
	mockPosts := []Post{
		{PostID: 3, PostClassID: 1, PostName: "GroupMe Link", PostContent: "www.groupme.com", PostVotes: 9},
	}

	mockPostController.
		EXPECT().
		GetClassPostsByName(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockPosts)
		})

	req, err := http.NewRequest("GET", "/getPostsByClassId/"+className, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"className": className})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockPostController.GetClassPostsByName)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response into a slice of Post
	var returnedPosts []Post
	err = json.Unmarshal(rr.Body.Bytes(), &returnedPosts)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the data
	if !reflect.DeepEqual(returnedPosts, mockPosts) {
		t.Errorf("handler returned unexpected data: got %v want %v", returnedPosts, mockPosts)
	}
}

// IncreaseVotes test
func TestIncreasePostVotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostController := controller.NewMockPostController(ctrl)

	postID := "4"
	mockPostController.
		EXPECT().
		IncreasePostVote(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Post with ID = " + postID + " was updated"))
		})

	req, err := http.NewRequest("PUT", "/increasePostVotes/"+postID, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"postID": postID})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockPostController.IncreasePostVote)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `Post with ID = 4 was updated`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDecreasePostVotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostController := controller.NewMockPostController(ctrl)

	postID := "1"
	mockPostController.
		EXPECT().
		DecreasePostVotes(gomock.Any(), gomock.Any()).
		DoAndReturn(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Post with ID = " + postID + " was updated"))
		})

	req, err := http.NewRequest("PUT", "/decreasePostVotes/"+postID, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"postID": postID})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockPostController.DecreasePostVotes)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `Post with ID = 1 was updated`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
