package main

import (
	"testing"
)

func TestUrl(t *testing.T) {
	InitDBForTest()
	email := "takuto.yoshikai.url@gmail.com"
	password := "takuto01"
	dstUrl := "https://yoshikai.net"
	_ = CreateUser(email, password)
	user := GetUser(email)
	url, errs := CreateURL(user, dstUrl)
	if errs != nil && len(errs) > 0 {
		t.Fatal(errs)
		return
	}
	if url == nil {
		t.Fatal("url does not exists")
		return
	}
	if len(url.SrcId) != 5 {
		t.Fatal("url id is not 5")
		return
	}
	if url.DstUrl != dstUrl {
		t.Fatal("dst url was not saved")
		return
	}
	if url.UserId != int(user.ID) {
		t.Fatal("user id was not saved")
		return
	}
}
