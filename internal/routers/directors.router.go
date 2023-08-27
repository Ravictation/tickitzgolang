package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/middleware"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func directors_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/directors")

	repo := repositories.NewDirectors(d)
	handler := handlers.NewDirectors(repo)

	route.GET("/", handler.Get_Directors)
	route.POST("/", middleware.Authjwt("admin"), handler.Post_Directors)
	route.PUT("/:id", middleware.Authjwt("admin"), handler.Put_Directors)
	route.DELETE("/:id", middleware.Authjwt("admin"), handler.Delete_Directors)

}
