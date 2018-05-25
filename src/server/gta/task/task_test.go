package task

import (
	"testing"
)

func TestAdd(t *testing.T) {
	id:=Add(
		0,
		"task modle implement",
		"implement task modle and test it")
	println("added Task is=",id)
}

func TestGetAll(t *testing.T){
	taskStr:=GetAll()
	println(taskStr)
}
