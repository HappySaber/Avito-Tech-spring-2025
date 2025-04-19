package routes

import (
	controllers "PVZ/src/controller"
	"PVZ/src/midlleware"

	"github.com/gin-gonic/gin"
)

func PVZRoutes(r *gin.Engine) {
	r.POST("/dummyLogin", controllers.DummyLogin)
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Signup)

	userGroup := r.Group("/pvz").Use(midlleware.IsAuthorized())
	{
		userGroup.POST("/createpvz", midlleware.IsModerator(), controllers.CreatePVZ)
	}
}
