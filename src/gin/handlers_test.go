package gin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	s "github.com/heroku/go-getting-started/src"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	exitVal := m.Run()
	os.Exit(exitVal)
}
func TestRegisterEmail(t *testing.T) {
	invalidEmail := "jonathan.seubert"
	duplicateEmail := "jonathan.seubert3@megajon.com"
	newEmail := "finisher1017@gmail.com"
	form := url.Values{}
	router := SetupRouter()

	test1 := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	form.Add("email", invalidEmail)
	req.PostForm = form
	router.ServeHTTP(test1, req)
	var invalidResponseObject s.Message
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
	var duplicateResponseObject s.Message
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
	var newEmailResponseObject s.Message
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
	invalidEmail := "jonathan.seubert"
	nonExistentEmail := "none.existent@megajon.com"
	newEmail := "finisher1017@gmail.com"
	form := url.Values{}
	router := SetupRouter()

	test1 := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/delete", strings.NewReader(form.Encode()))
	form.Add("email", invalidEmail)
	req.PostForm = form
	router.ServeHTTP(test1, req)
	var invalidResponseObject s.Message
	err := json.Unmarshal(test1.Body.Bytes(), &invalidResponseObject)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Printf("invalid response message: %v", invalidResponseObject.Message)
	if invalidResponseObject.Message != "invalid email" {
		t.Fatal("invalid email was allowed")
	}

	test2 := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/delete", strings.NewReader(form.Encode()))
	form.Del("email")
	form.Add("email", nonExistentEmail)
	req.PostForm = form
	router.ServeHTTP(test2, req)
	var nonExistentEmailObject s.Message
	err = json.Unmarshal(test2.Body.Bytes(), &nonExistentEmailObject)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Printf("non existent response message: %v", nonExistentEmailObject.Message)
	if nonExistentEmailObject.Message != "no email found" {
		t.Fatal("email already exists")
	}

	test3 := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/delete", strings.NewReader(form.Encode()))
	form.Del("email")
	form.Add("email", newEmail)
	req.PostForm = form
	router.ServeHTTP(test3, req)
	var newEmailObject s.Message
	err = json.Unmarshal(test3.Body.Bytes(), &newEmailObject)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Printf("new email response message: %v", newEmailObject)
	if newEmailObject.Message != "email deleted" {
		t.Fatal("unable to delete email")
	}
}
