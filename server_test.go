package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func SendJSON(r *gin.Engine, path string, bodyString string) (int, string) {
	req := httptest.NewRequest("POST", path, bytes.NewBuffer([]byte(bodyString)))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	return rec.Code, string(body)
}

func GetRequest(r *gin.Engine, path string) (int, string) {
	req := httptest.NewRequest("GET", path, nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	return rec.Code, string(body)
}

func TestServer(t *testing.T) {
	err := godotenv.Load("env.test")
	if err != nil {
		t.Fatal(err)
		return
	}
	InitDBForTest()
	r := InitServer()
	statusCode, responseText := SendJSON(r, "/create", `{"url":"https://yoshikai.net"}`)
	if statusCode != 200 {
		t.Fatal("status code is not 200")
		return
	}
	if !strings.HasPrefix(responseText, "http://localhost:8080/") {
		t.Fatal("it is not a url" + responseText)
		return
	}

	statusCode, _ = GetRequest(r, responseText)
	if statusCode != 302 {
		t.Fatal("didn't redirect")
		return
	}

	statusCode, _ = SendJSON(r, "/user", `{"email": "takuto.yoshikai.post@gmail.com", "password": "takuto01"}`)
	if statusCode != 201 {
		t.Fatal("couldn't create a user")
		return
	}
}
