package routers

import (
	"github.com/vaisakhns/ecommerse-vs/controllers"

	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRoutes *gin.Engine)  {
	
	incomingRoutes.POST("/user/signup", controllers.SignUp())
    incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.SearchProducts())
	incomingRoutes.GET("/users/search",controllers.SearchProductByQuery())
}