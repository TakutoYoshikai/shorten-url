package main

import (
	"github.com/gin-gonic/gin"
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

func InitServer() {
	r := gin.Default()
	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		url := GetURL(id)
		if url == nil {
			c.String(404, "")
			return
		}
		c.Redirect(302, url.DstUrl)
	})
	r.POST("/create", func(c *gin.Context) {
		var body URLRequestBody
		c.BindJSON(&body)
		url, errs := CreateURL(nil, body.Url)
		if url == nil || (errs != nil && len(errs) > 0) {
			c.String(400, "")
			return
		}
		c.String(200, CreateURLString(url.SrcId))
	})
	r.POST("/user", func(c *gin.Context) {
		var body CreateUserRequestBody
		c.BindJSON(&body)
		user, errs := CreateUser(body.Email, body.Password)
		if user == nil || errs != nil && len(errs) > 0 {
			c.String(400, "")
			return
		}
		c.String(201, "")
	})
	r.GET("/login", func(c *gin.Context) {
		var body LoginRequestBody
		c.BindJSON(&body)
		user := Login(body.Email, body.Password)
		if user == nil {
			c.String(401, "")
			return
		}
	})
	r.Run()
}
