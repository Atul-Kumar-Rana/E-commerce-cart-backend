package routes

import(
	"github.com/Atul-Kumar-Rana/E-commerce-cart-backend/controllers"
	"github.com/gin-gonic/gin"
)
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup",controllers.signup())
	incomingRoutes.POST("/users/login",controllers.login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.searchProduct())
	incomingRoutes.GET("/users/search",controllers.searchProductByQuery())
}