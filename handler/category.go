package handler

import (
	"final_project_3/dto"
	"final_project_3/pkg/errs"
	"final_project_3/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) categoryHandler {
	return categoryHandler{categoryService: categoryService}
}

func (ch *categoryHandler) CreateCategory(ctx *gin.Context) {
	var categoryPayload dto.NewCategoryRequest
	if err := ctx.ShouldBindJSON(&categoryPayload); err != nil {
		errBind := errs.NewUnprocessableEntityResponse("invalid request body")
		ctx.JSON(errBind.Status(), errBind)
		return
	}
	response, err := ch.categoryService.CreateCategory(&categoryPayload)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (ch *categoryHandler) GetCategories(ctx *gin.Context) {
	response, err := ch.categoryService.GetCategoriesWithTasks()
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryHandler) UpdateCategory(ctx *gin.Context) {
	var categoryPayload dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&categoryPayload); err != nil {
		errBind := errs.NewUnprocessableEntityResponse("invalid request body")
		ctx.JSON(errBind.Status(), errBind)
		return
	}
	param := ctx.Param("categoryId")
	categoryId, errConv := strconv.Atoi(param)
	if errConv != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errConv.Error(),
		})
		return
	}
	categoryPayload.ID = uint(categoryId)
	response, err := ch.categoryService.UpdateCategory(&categoryPayload)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryHandler) DeleteCategory(ctx *gin.Context) {
	param := ctx.Param("categoryId")
	categoryId, errConv := strconv.Atoi(param)
	if errConv != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errConv.Error(),
		})
		return
	}
	err := ch.categoryService.DeleteCategory(categoryId)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}
