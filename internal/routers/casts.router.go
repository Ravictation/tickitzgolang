package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func casts_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/casts")

	repo := repositories.NewCasts(d)
	handler := handlers.NewCasts(repo)

	route.GET("/", handler.Get_Casts)
	route.POST("/", handler.Post_Casts)
	route.PUT("/:id", handler.Put_Casts)
	route.DELETE("/:id", handler.Delete_Casts)

}
