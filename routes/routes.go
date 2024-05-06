package routes

import(
	"github.com/yasirnabil534/eCommerceProject/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/addmin/add-product", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/product-view", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}