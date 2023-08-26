package pkg

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Server(router *gin.Engine) *http.Server {
	var addr string = "0.0.0.0:8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	corss := cors.AllowAll()
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      corss.Handler(router),
	}

	return srv
}
