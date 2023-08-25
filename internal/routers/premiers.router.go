package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/middleware"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func premiers_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/premiers")

	repo := repositories.NewPremiers(d)
	handler := handlers.NewPremiers(repo)

	route.GET("/", handler.Get_Premiers)
	route.POST("/", middleware.UploadFile("image"), handler.Post_Premiers)
	route.PUT("/:id", middleware.UploadFile("image"), handler.Put_Premiers)
	route.DELETE("/:id", handler.Delete_Premiers)

}
