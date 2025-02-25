package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahfi/to-do-list-app/dto"
	"github.com/kahfi/to-do-list-app/helpers"
	"github.com/kahfi/to-do-list-app/models"
	"github.com/kahfi/to-do-list-app/repository"
)

type ToDoListController struct{}

// GET DATA
func (ctrl *ToDoListController) GetData(ctx *gin.Context) {
	ownerId := ctx.GetHeader("Authorization")
	var response []dto.ToDoListResponse
	var toDoGroup []models.ToDoGroup
	if err := repository.GetData(&toDoGroup, ownerId); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  true,
			Message: "Gagal Mendapatkan Data",
		})
		return
	}
	for _, data := range toDoGroup {
		response = append(response, dto.ToDoListResponse{
			Id:        data.Id,
			GroupName: data.GroupName,
			Details:   data.ToDoList,
			Status:    data.Status,
			OwnerId:   data.OwnerId,
		})
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    response,
		Message: "Success",
	})
}

// GET DATA
func (ctrl *ToDoListController) GetDetails(ctx *gin.Context) {
	groupId := ctx.Param("groupid")
	var toDoList []models.ToDoList
	if err := repository.GetDataByPredicate(&toDoList, "groupid", groupId); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  true,
			Message: "Gagal Mendapatkan Data",
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    toDoList,
		Message: "Success",
	})
}

// CREATE DATA
func (ctl *ToDoListController) CreateList(ctx *gin.Context) {
	var body dto.ToDoListPostRequest
	if err := ctx.BindJSON(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	var toDoList models.ToDoList
	toDoList.ToDoTask = body.ToDoTask
	toDoList.Status = body.Status
	toDoList.GroupId = body.GroupId
	toDoList.OwnerId = body.OwnerId
	if err := repository.Save(&toDoList); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  true,
			Message: "Gagal Membuat Data",
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
	})
}

// UPDATE DATA
func (ctrl *ToDoListController) UpdateData(ctx *gin.Context) {
	var body dto.ToDoListPutRequest
	if err := ctx.BindJSON(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	var todoList models.ToDoList
	todoList.ToDoTask = body.ToDoTask
	todoList.GroupId = body.GroupId
	if err := repository.UpdateData(&todoList, body.Id); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  true,
			Message: "Gagal Mengubah data",
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
	})

}

// CHECK UNCHECK DATA
func (ctrl *ToDoListController) CheckData(ctx *gin.Context) {
	var body dto.ToDoListCheck
	if err := ctx.BindJSON(&body); err != nil {
		webResponse := helpers.Response{
			Code:    http.StatusBadRequest,
			Status:  true,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	var todoList models.ToDoList
	todoList.Status = body.Status
	if err := repository.CheckUnchekData(&todoList, body.Id); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  true,
			Message: "Gagal Mengubah data",
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
	})
}

func (ctrl *ToDoListController) DeleteData(ctx *gin.Context) {
	id := ctx.Param("id")
	var todoList models.ToDoList
	if err := repository.SoftDeleteById(&todoList, "id", id); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Status:  true,
			Message: "Gagal Menghapus data",
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
	})
}
