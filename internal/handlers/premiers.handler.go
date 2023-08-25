package handlers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/repositories"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Premiers struct {
	*repositories.Repo_Premiers
}

func NewPremiers(r *repositories.Repo_Premiers) *Handler_Premiers {
	return &Handler_Premiers{r}
}

func (h *Handler_Premiers) Get_Premiers(ctx *gin.Context) {
	var premiers models.Premiers

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&premiers); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&premiers, page, limit)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}

func (h *Handler_Premiers) Post_Premiers(ctx *gin.Context) {
	var premiers models.Premiers

	if err := ctx.ShouldBind(&premiers); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	premiers.Image = ctx.MustGet("image").(string)

	var err_val error
	_, err_val = govalidator.ValidateStruct(&premiers)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Insert_Data(&premiers)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Premiers) Put_Premiers(ctx *gin.Context) {
	var premiers models.Premiers
	premiers.Id_premier = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(premiers.Id_premier)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	premiers.Image = ctx.MustGet("image").(string)

	if err := ctx.ShouldBind(&premiers); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	var err_val error
	_, err_val = govalidator.ValidateStruct(&premiers)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Update_Data(&premiers)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Premiers) Delete_Premiers(ctx *gin.Context) {
	var premiers models.Premiers
	premiers.Id_premier = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(premiers.Id_premier)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&premiers); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&premiers)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
