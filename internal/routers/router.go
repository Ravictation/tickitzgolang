package routers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(config.CorsConfig))

	user(router, db)
	auth(router, db)
	return router
}
