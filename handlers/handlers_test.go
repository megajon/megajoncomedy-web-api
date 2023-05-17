package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	u "github.com/heroku/go-getting-started/helpers"
)

// func TestMain(m *testing.M) {
// 	db := u.Connect()
// 	exitVal := m.Run()
// 	os.Exit(exitVal)
// }

func TestRegisterEmail(t *testing.T) {
	invalidEmail := "jonathan.seubert"
	duplicateEmail := "jonathan.seubert@megajon.com"
	newEmail := "jonathan.seubert3@megajon.com"
	form := url.Values{}
	router := SetupRouter()

	test1 := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	form.Add("email", invalidEmail)
	req.PostForm = form
	router.ServeHTTP(test1, req)
	var invalidResponseObject u.Message
	err := json.Unmarshal(test1.Body.Bytes(), &invalidResponseObject)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Printf("invalid response message: %v", invalidResponseObject.Message)
	if invalidResponseObject.Message != "invalid email" {
		t.Fatal("invalid email was allowed")
	}

	test2 := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	form.Del("email")
	form.Add("email", duplicateEmail)
	req.PostForm = form
	router.ServeHTTP(test2, req)
	var duplicateResponseObject u.Message
	err = json.Unmarshal(test2.Body.Bytes(), &duplicateResponseObject)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Printf("duplicate response message: %v", duplicateResponseObject.Message)
	if duplicateResponseObject.Message != "database error" {
		t.Fatal("duplicate email was allowed")
	}

	test3 := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	form.Del("email")
	form.Add("email", newEmail)
	req.PostForm = form
	router.ServeHTTP(test3, req)
	var newEmailResponseObject u.Message
	err = json.Unmarshal(test3.Body.Bytes(), &newEmailResponseObject)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Printf("new email response message: %v", newEmailResponseObject.Message)
	if newEmailResponseObject.Message != "success" {
		t.Fatal("store email failed")
	}
}

func TestDeleteEmail(t *testing.T) {

}
