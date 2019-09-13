package main

import (
	"testing"
)

func TestInitDB(t *testing.T){
	err := initDB("mysql", "root:password@tcp(127.0.0.1:8081)/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
			t.Errorf("initDB fail, expect nil got %v",err)
	}
}