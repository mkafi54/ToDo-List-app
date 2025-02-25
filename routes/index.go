package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	routes "github.com/kahfi/to-do-list-app/routes/group"
)

func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	routes.ToDoListRoutes(route)
}
