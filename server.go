package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type URLRequestBody struct {
	Url string `json:"url"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateURLString(id string) string {
	protocol := os.Getenv("protocol")
	host := os.Getenv("host")
	return protocol + "://" + host + "/" + id
}

func RedirectByIdRequest(c *gin.Context) {
	id := c.Param("id")
	url := GetURL(id)
	if url == nil {
		c.String(404, "")
		return
	}
	c.Redirect(302, url.DstUrl)
}

func CreateURLRequest(c *gin.Context) {
	var body URLRequestBody
	c.BindJSON(&body)
	url, errs := CreateURL(nil, body.Url)
	if url == nil || (errs != nil && len(errs) > 0) {
		c.String(400, "")
		return
	}
	c.String(200, CreateURLString(url.SrcId))
}

func CreateUserRequest(c *gin.Context) {
	var body CreateUserRequestBody
	c.BindJSON(&body)
	user, errs := CreateUser(body.Email, body.Password)
	if user == nil || errs != nil && len(errs) > 0 {
		c.String(400, "")
		return
	}
	c.String(201, "")
}

func LoginRequest(c *gin.Context) {
	var body LoginRequestBody
	c.BindJSON(&body)
	user := Login(body.Email, body.Password)
	if user == nil {
		c.String(401, "")
		return
	}
}

func InitServer() *gin.Engine {
	file, err := os.OpenFile("logs/gin.log", os.O_WRONLY|os.O_CREATE, 0666)
	if file != nil && err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}
	r := gin.Default()
	r.GET("/:id", RedirectByIdRequest)
	r.POST("/create", CreateURLRequest)
	r.POST("/user", CreateUserRequest)
	r.POST("/login", LoginRequest)
	return r
}
