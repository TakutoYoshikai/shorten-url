package main

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	InitDBForTest()
	err := CreateUser("takuto.yoshikai@gmail.com", "takuto01")
	if err != nil && len(err) > 0 {
		t.Fatal(err)
	}
	err = CreateUser("helloworld", "takuto01")
	if err == nil || len(err) == 0 {
		t.Fatal("It could register not email string as email")
	}
}
