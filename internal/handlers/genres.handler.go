package handlers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/repositories"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Genres struct {
	*repositories.Repo_Genres
}

func NewGenres(r *repositories.Repo_Genres) *Handler_Genres {
	return &Handler_Genres{r}
}

func (h *Handler_Genres) Get_Genres(ctx *gin.Context) {
	var genres models.Genres

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&genres); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&genres, page, limit)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}

func (h *Handler_Genres) Post_Genres(ctx *gin.Context) {
	var genres models.Genres

	if err := ctx.ShouldBind(&genres); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	var err_val error
	_, err_val = govalidator.ValidateStruct(&genres)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Insert_Data(&genres)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Genres) Put_Genres(ctx *gin.Context) {
	var genres models.Genres
	genres.Id_genre = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(genres.Id_genre)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&genres); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	var err_val error
	_, err_val = govalidator.ValidateStruct(&genres)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Update_Data(&genres)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Genres) Delete_Genres(ctx *gin.Context) {
	var genres models.Genres
	genres.Id_genre = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(genres.Id_genre)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&genres); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&genres)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
