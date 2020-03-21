package main

import (
	"github.com/gin-gonic/gin"
)

type URLRequestBody struct {
	Url string `json:"url"`
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
		c.String(200, url.SrcId)
	})
	r.Run()
}
