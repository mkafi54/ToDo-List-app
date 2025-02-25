package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kahfi/to-do-list-app/controllers"
)

func ToDoListRoutes(route *gin.Engine) {
	var ctrl controllers.ToDoListController

	route.GET("/to-do-list/get", ctrl.GetData)
	route.GET("/to-do-list/get-details/:groupid", ctrl.GetDetails)
	route.POST("/to-do-list/create", ctrl.CreateList)
	route.PUT("/to-do-list/update", ctrl.UpdateData)
	route.PUT("/to-do-list/check-uncheck", ctrl.CheckData)
	route.DELETE("/to-do-list/delete", ctrl.DeleteData)

}
