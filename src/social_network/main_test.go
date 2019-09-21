package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"bytes"
	"io/ioutil"
)

type (
	testUserModel struct {
		ID        uint
		Name     string `json:"name"`
  		Email    string    `json:"email"`
	}
)


func TestInitDB(t *testing.T){
	err := initDB("mysql", "root:password@tcp(127.0.0.1:8081)/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
			t.Errorf("initDB fail, expect nil got %v",err)
	}
}

func TestCreate(t *testing.T){
	//var testData = []byte( `{{name:"name1", email:"email1@mail.com"}, {name:"name2", email:"email2@mail.com"}}`)
	initDB("mysql", "root:password@tcp(127.0.0.1:8081)/mysql?charset=utf8&parseTime=True&loc=Local")
	var testData = []byte( `{"name":"name1", "email":"email1@mail.com"}`)
	var testUser userModel
	//var resultUser userModel

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	//create new request
	req, err1 := http.NewRequest("", "", bytes.NewBuffer(testData))
	if err1 != nil {
		t.Fatal(err1)
	}
	//add request to context
	c.Request = req

	_, err2 := testUser.create(c)
	//println(returnUser)
	if err2 != nil{
			t.Errorf("Create user fail, expect nil got %v",err2)
	}	

	//read from db & compare
	/*
	db.First(&resultUser, returnUser.ID)
	t.Errorf("Create user fail, expect %v got %v",returnUser, resultUser)
	if testUser != resultUser {
		t.Errorf("Create user fail, expect %v got %v",returnUser, resultUser)
	}
	*/
}

func TestRead(t *testing.T){
	initDB("mysql", "root:password@tcp(127.0.0.1:8081)/mysql?charset=utf8&parseTime=True&loc=Local")
	var testUser userModel
	var id = "4"
	resultUser, err2 := testUser.read(id)
	//println(returnUser)
	if err2 != nil{
			t.Errorf("Read user fail, expect nil got %v",err2)
	}

	if resultUser.Name == "" || resultUser.Email == "" {
		t.Errorf("Read user fail, expect Name and Email got nothing, resultUser: %v", resultUser)
	}		
}

func TestUpdate(t *testing.T){
	initDB("mysql", "root:password@tcp(127.0.0.1:8081)/mysql?charset=utf8&parseTime=True&loc=Local")
	var testData = []byte( `{"name":"updateName1", "email":"email1@mail.com"}`)
	var testUser userModel
	//var resultUser userModel

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	//create new request
	req, err1 := http.NewRequest("", "", bytes.NewBuffer(testData))
	if err1 != nil {
		t.Fatal(err1)
	}
	//add request to context
	c.Request = req

	_, err2 := testUser.update("2",c)
	if err2 != nil{
			t.Errorf("Update user fail, expect nil got %v",err2)
	}	

	//read from db & compare
	/*
	db.First(&resultUser, returnUser.ID)
	t.Errorf("Create user fail, expect %v got %v",returnUser, resultUser)
	if testUser != resultUser {
		t.Errorf("Create user fail, expect %v got %v",returnUser, resultUser)
	}
	*/
}

func (user testUserModel) create(c *gin.Context) (returnUser userModel, err error){
	returnUser.ID = "1"
    return returnUser, nil
}

func (user testUserModel) read(id string) (returnUser userModel, err error){
	returnUser.Name = "testName"
	returnUser.Email = "testEmail"
    return returnUser, nil
}

func (user testUserModel) update(id string, c *gin.Context) (returnUser userModel, err error){
	returnUser.ID = "2"
    return returnUser, nil
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, nil)
   w := httptest.NewRecorder()
   r.ServeHTTP(w, req)
   return w
}

func TestCreateUser(t *testing.T){
	var testUser testUserModel
	router := setRoute(testUser)
	w := performRequest(router, "POST", "api/vi/users/")
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetSingleUser(t *testing.T){
	var testUser testUserModel
	router := setRoute(testUser)
	w := performRequest(router, "GET", "api/vi/users/1")
	body, _ := ioutil.ReadAll(w.Body)
	result := string(body)
	var chk = "{\"data\":{\"ID\":\"\",\"name\":\"testName\",\"email\":\"testEmail\"},\"status\":200}\n"
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// need implement better response check in the future
	if result != chk {
		t.Errorf("handler returned wrong data: got %v want %v", result, chk)
	}
}

func TestUpdateUser(t *testing.T){
	var testUser testUserModel
	router := setRoute(testUser)
	w := performRequest(router, "PATCH", "api/vi/users/2")
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}