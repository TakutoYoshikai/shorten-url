package main

import (
	"testing"
)

func TestUser(t *testing.T) {
	InitDBForTest()
	email := "takuto.yoshikai@gmail.com"
	password := "takuto01"
	user, errs := CreateUser(email, password)
	if errs != nil && len(errs) > 0 {
		t.Fatal(errs)
		return
	}
	if user == nil {
		t.Fatal("user isn't created")
		return
	}
	user, errs = CreateUser("helloworld", password)
	if errs == nil || len(errs) == 0 || user != nil {
		t.Fatal("It could register not email string as email")
		return
	}
	user = GetUser(email)
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
