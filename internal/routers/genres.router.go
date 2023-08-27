package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/middleware"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func genres_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/genres")

	repo := repositories.NewGenres(d)
	handler := handlers.NewGenres(repo)

	route.GET("/", handler.Get_Genres)
	route.POST("/", middleware.Authjwt("admin"), handler.Post_Genres)
	route.PUT("/:id", middleware.Authjwt("admin"), handler.Put_Genres)
	route.DELETE("/:id", middleware.Authjwt("admin"), handler.Delete_Genres)

}
