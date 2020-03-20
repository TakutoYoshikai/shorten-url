package main

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	InitDBForTest()
	err := CreateUser("takuto.yoshikai@gmail.com", "takuto01")
	if err != nil && len(err) > 0 {
		t.Fatal(err)
		return
	}
	err = CreateUser("helloworld", "takuto01")
	if err == nil || len(err) == 0 {
		t.Fatal("It could register not email string as email")
		return
	}
	user := GetUser("takuto.yoshikai@gmail.com")
	if user == nil {
		t.Fatal("Couldn't get user")
		return
	}
	if user.Email != "takuto.yoshikai@gmail.com" {
		t.Fatal("Got wrong user")
		return
	}
	user = GetUser("notexists")
	if user != nil {
		t.Fatal("got not exists user")
		return
	}
}
