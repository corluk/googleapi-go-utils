package server

import (
	"fmt"

	"github.com/corluk/googleapi-go-utils/auth"
	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {

	r.GET("/api/google/login/url", func(c *gin.Context) {
		var requestLoginUrl auth.RequestLoginUrl

		err := c.BindJSON(&requestLoginUrl)
		if err != nil {
			c.String(400, fmt.Sprintf("message %s", err))
			return
		}

		url, err := requestLoginUrl.GetUrl()
		if err != nil {
			c.String(401, fmt.Sprintf("message %s", err))
			return
		}

		c.String(200, url.String())

	})

}
