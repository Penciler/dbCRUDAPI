package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"bytes"
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