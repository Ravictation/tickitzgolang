package routers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(config.CorsConfig))
	router.Use(middleware.CORSMiddleware)

	user(router, db)
	auth(router, db)
	directors_router(router, db)
	casts_router(router, db)
	genres_router(router, db)
	movies_router(router, db)
	premiers_router(router, db)
	times_schedules_router(router, db)

	return router
}
