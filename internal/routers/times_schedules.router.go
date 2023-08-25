package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func times_schedules_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/times_schedules")

	repo := repositories.NewTimesSchedules(d)
	handler := handlers.NewTimesSchedules(repo)

	route.GET("/:id", handler.Get_Times_Schedules_by_Id)
	route.GET("/", handler.Get_Times_Schedules)
	// route.POST("/", handler.Post_Times_Schedules)
	// route.PUT("/:id", handler.Put_Times_Schedules)
	// route.DELETE("/:id", handler.Delete_Times_Schedules)

}
