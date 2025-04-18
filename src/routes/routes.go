package routes

import (
	controllers "PVZ/src/controller"

	"github.com/gin-gonic/gin"
)

func PVZRoutes(r *gin.Engine) {
	r.POST("/dummyLogin", controllers.DummyLogin)
}
