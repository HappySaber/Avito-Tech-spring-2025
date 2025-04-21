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
		userGroup.POST("/create-pvz", midlleware.IsModerator(), controllers.CreatePVZ)
		userGroup.POST("/:pvzid/initiate-reception", midlleware.IsPVZemployee(), controllers.InitiateReceivingHandler)
		userGroup.POST("/:pvzid/add-product", midlleware.IsPVZemployee(), controllers.AddProductHandler)
		userGroup.DELETE("/:pvzid/delete-last-product", midlleware.IsPVZemployee(), controllers.DeleteLastProduct)
		userGroup.POST("/:pvzid/close-reception", midlleware.IsPVZemployee(), controllers.CloseReception)
		userGroup.GET("/data", midlleware.IsPVZemployeeOrModerator(), controllers.GetPVZDataHandler)
	}
}
