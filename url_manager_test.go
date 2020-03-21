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
	tmpUrl := GetURL(url.SrcId)
	if tmpUrl == nil {
		t.Fatal("url not found")
		return
	}
	if tmpUrl.SrcId != url.SrcId || tmpUrl.DstUrl != url.DstUrl || tmpUrl.ID != url.ID || tmpUrl.UserId != url.UserId {
		t.Fatal("url property is wrong")
		return
	}
	url = GetURL("notexists")
	if url != nil {
		t.Fatal("found not existing url")
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
	tmpUrl = GetURL(url.SrcId)
	if tmpUrl == nil {
		t.Fatal("url not found")
		return
	}
	if tmpUrl.SrcId != url.SrcId || tmpUrl.DstUrl != url.DstUrl || tmpUrl.ID != url.ID || tmpUrl.UserId != url.UserId {
		t.Fatal("url property is wrong")
		return
	}
	err := DeleteURL(url)
	if err != nil {
		t.Fatal("couldn't delete url")
		return
	}
	tmpUrl = GetURL(url.SrcId)
	if tmpUrl != nil {
		t.Fatal("couldn't delete url")
		return
	}

}
