package handlers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/repositories"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Directors struct {
	*repositories.Repo_Directors
}

func NewDirectors(r *repositories.Repo_Directors) *Handler_Directors {
	return &Handler_Directors{r}
}

func (h *Handler_Directors) Get_Directors(ctx *gin.Context) {
	var director models.Directors

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&director); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&director, page, limit)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}

func (h *Handler_Directors) Post_Directors(ctx *gin.Context) {
	var director models.Directors

	if err := ctx.ShouldBind(&director); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	var err_val error
	_, err_val = govalidator.ValidateStruct(&director)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Insert_Data(&director)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Directors) Put_Directors(ctx *gin.Context) {
	var director models.Directors
	director.Id_director = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(director.Id_director)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&director); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	var err_val error
	_, err_val = govalidator.ValidateStruct(&director)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Update_Data(&director)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Directors) Delete_Directors(ctx *gin.Context) {
	var director models.Directors
	director.Id_director = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(director.Id_director)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&director); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&director)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
