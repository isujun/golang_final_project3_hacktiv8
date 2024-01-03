package handler

import (
	"final_project_3/dto"
	"final_project_3/pkg/errs"
	"final_project_3/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) taskHandler {
	return taskHandler{taskService: taskService}
}

func (th *taskHandler) CreateTask(ctx *gin.Context) {
	var taskPayload dto.NewTaskRequest
	if err := ctx.ShouldBindJSON(&taskPayload); err != nil {
		errBind := errs.NewUnprocessableEntityResponse("invalid request body")
		ctx.JSON(errBind.Status(), errBind)
		return
	}
	jwtClaims := ctx.MustGet("user").(jwt.MapClaims)
	taskPayload.UserID = uint(jwtClaims["id"].(float64))
	response, err := th.taskService.CreateTask(&taskPayload)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (th *taskHandler) GetTasks(ctx *gin.Context) {
	response, err := th.taskService.GetTasks()
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (th taskHandler) UpdateTask(ctx *gin.Context) {
	var taskPayload dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&taskPayload); err != nil {
		errBind := errs.NewUnprocessableEntityResponse("invalid request body")
		ctx.JSON(errBind.Status(), errBind)
		return
	}
	param := ctx.Param("taskId")
	taskId, errConv := strconv.Atoi(param)
	if errConv != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errConv.Error(),
		})
		return
	}
	taskPayload.Title = strconv.Itoa(taskId) 
	response, err := th.taskService.UpdateTask(&taskPayload)
	if err != nil  {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (th *taskHandler) UpdateStatus(ctx *gin.Context) {
	var taskPayload dto.UpdateStatusResquest
	if err := ctx.ShouldBindJSON(&taskPayload); err != nil {
		errBinJson := errs.NewUnprocessableEntityResponse("invalid json request body")
		ctx.JSON(errBinJson.Status(), errBinJson)
		return
	}
	param := ctx.Param("taskId")
	taskId, errConv := strconv.ParseBool(param)
	if errConv != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errConv.Error(),
		})
		return
	}
	taskPayload.Status = bool(taskId)
	response, err := th.taskService.UpdateStatus(&taskPayload)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
