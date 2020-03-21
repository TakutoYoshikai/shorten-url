package main

import (
	"testing"
)

func TestUser(t *testing.T) {
	InitDBForTest()
	email := "takuto.yoshikai@gmail.com"
	password := "takuto01"
	err := CreateUser(email, password)
	if err != nil && len(err) > 0 {
		t.Fatal(err)
		return
	}
	err = CreateUser("helloworld", password)
	if err == nil || len(err) == 0 {
		t.Fatal("It could register not email string as email")
		return
	}
	user := GetUser(email)
	if user == nil {
		t.Fatal("Couldn't get user")
		return
	}
	if user.Email != email {
		t.Fatal("Got wrong user")
		return
	}
	user = GetUser("notexists")
	if user != nil {
		t.Fatal("got not exists user")
		return
	}
	user = Login(email, password)
	if user == nil {
		t.Fatal("couldn't login")
		return
	}
	user = Login(email, "wrongpass")
	if user != nil {
		t.Fatal("logged in by wrong password")
		return
	}
	user = Login("notexists", password)
	if user != nil {
		t.Fatal("logged in as not existing user")
		return
	}

}
