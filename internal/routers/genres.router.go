package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func genres_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/genres")

	repo := repositories.NewGenres(d)
	handler := handlers.NewGenres(repo)

	route.GET("/", handler.Get_Genres)
	route.POST("/", handler.Post_Genres)
	route.PUT("/:id", handler.Put_Genres)
	route.DELETE("/:id", handler.Delete_Genres)

}
