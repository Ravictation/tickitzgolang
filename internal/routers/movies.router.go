package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/middleware"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func movies_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/movies")

	repo := repositories.NewMovies(d)
	handler := handlers.NewMovies(repo)

	route.GET("/", handler.Get_Movies)
	route.POST("/", middleware.UploadFile("image"), middleware.UploadFile2("cover_image"), handler.Post_Movies)
	route.PUT("/:id", middleware.UploadFile("image"), middleware.UploadFile2("cover_image"), handler.Put_Movies)
	route.DELETE("/:id", handler.Delete_Movies)

}
