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

	route.GET("/:id", handler.Get_Movies_by_Id)
	route.GET("/", handler.Get_Movies)
	route.POST("/", middleware.Authjwt("admin"), middleware.UploadFile("image"), middleware.UploadFile2("cover_image"), handler.Post_Movies)
	route.PUT("/:id", middleware.Authjwt("admin"), middleware.UploadFile("image"), middleware.UploadFile2("cover_image"), handler.Put_Movies)
	route.DELETE("/:id", middleware.Authjwt("admin"), handler.Delete_Movies)

}
