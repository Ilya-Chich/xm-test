package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	v1 "xm-test-ilya-chicherin/internal/controllers/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	{
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}
	}
	router := gin.New()
	var err error
	if err = v1.New(router); err != nil {
		log.Fatal("error loading controllers")
	}
	if os.Getenv("DEBUG") == "true" {
		_ = router.Run(":" + os.Getenv("HTTP_PORT"))
	} else {
		_ = router.RunTLS(":"+os.Getenv("HTTP_PORT"), os.Getenv("SSL_CERT_PATH"), os.Getenv("SSL_KEY_PATH"))
	}
}
