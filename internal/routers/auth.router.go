package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/auth")

	repo := repositories.NewUser(d)
	handler := handlers.NewAuth(repo)

	route.POST("/login", handler.Login)

}
