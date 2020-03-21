package main

import (
	"testing"
)

func TestUrl(t *testing.T) {
	InitDBForTest()
	email := "takuto.yoshikai.url@gmail.com"
	password := "takuto01"
	dstUrl := "https://yoshikai.net"
	user, errs := CreateUser(email, password)
	if user == nil || (errs != nil && len(errs) > 0) {
		t.Fatal("user is not created")
		return
	}
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
	url, errs = CreateURL(nil, dstUrl)
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
	if url.UserId != -1 {
		t.Fatal("user id was not saved correctly")
		return
	}

}
