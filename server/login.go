package server

import (
	"fmt"
	"time"

	"github.com/corluk/googleapi-go-utils/auth"
	"github.com/gin-gonic/gin"
)

func GetUrl(r *gin.Engine) {

	r.GET("/api/google/auth/geturl/", func(c *gin.Context) {
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
		responseCode := new(ResponseGetUrl)
		responseCode.Time = time.Now()
		responseCode.URL = url.String()
		c.JSON(200, &responseCode)

	})

}

func ExchangeCode(r *gin.Engine) {

	r.GET("/api/google/auth/exchangecode/", func(c *gin.Context) {

	})
}

func RefreshToken(r *gin.Engine){
	r.GET("/api/google/auth/refreshcode/", func(c *gin.Context) {

		
	})
}