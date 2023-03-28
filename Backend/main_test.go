// main_test.go
package main

import (
	"Classical/Backend/controller"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
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

	expected := `[{"id":1,"className":"COP5000"},{"id":2,"className":"COP3502"},{"id":3,"className":"COP3503"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateClass(t *testing.T) {
	var jsonStr = []byte(`{"className":"COP1234"}`)

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
	expected := `{"id":6,"className":"COP1234"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

}
