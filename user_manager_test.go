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

func TestAllURL(t *testing.T) {
	InitDBForTest()
	email := "takuto.yoshikai.allurl@gmail.com"
	password := "takuto01"
	urlString := "https://yoshikai.net"
	user, errs := CreateUser(email, password)
	if errs != nil && len(errs) > 0 {
		t.Fatal(errs)
		return
	}
	if user == nil {
		t.Fatal("user isn't created")
		return
	}
	urls := AllURL(user)
	if urls != nil && len(urls) != 0 {
		t.Fatal("is not empty array of url")
		return
	}
	url, errs := CreateURL(user, urlString)
	if errs != nil && len(errs) > 0 {
		t.Fatal(errs)
		return
	}
	if url == nil {
		t.Fatal("url is not created")
		return
	}
	urls = AllURL(user)
	if len(urls) != 1 {
		t.Fatal("url is not created")
		return
	}
	if urls[0].ID != url.ID || urls[0].SrcId != url.SrcId || urls[0].DstUrl != url.DstUrl || urls[0].UserId != url.UserId {
		t.Fatal("url's property is wrong")
		return
	}
}
