package main

import (
	"github.com/corluk/googleapi-go-utils/server"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	server.GetUrl(r)
	
	r.Run()
}
