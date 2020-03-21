package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func SendJSON(r *gin.Engine, bodyString string) (int, string) {
	req := httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte(bodyString)))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	return rec.Code, string(body)
}

func TestServer(t *testing.T) {
	InitDB()
	r := InitServer()
	statusCode, responseText := SendJSON(r, `{"url":"https://yoshikai.net"}`)
	if statusCode != 200 {
		t.Fatal("status code is not 200")
		return
	}
	if !strings.HasPrefix(responseText, "http://localhost:8080/") {
		t.Fatal("it is not a url")
		return
	}
}
