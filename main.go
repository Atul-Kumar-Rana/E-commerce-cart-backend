package main

import (
	"log"
	"os"

	"github.com/Atul-Kumar-Rana/E-commerce-cart-backend/controllers"
	"github.com/Atul-Kumar-Rana/E-commerce-cart-backend/database"
	"github.com/Atul-Kumar-Rana/E-commerce-cart-backend/middleware"
	"github.com/Atul-Kumar-Rana/E-commerce-cart-backend/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port :=os.Getenv("PORT")
	if port ==""{
		port ="8080"
	}
	app := controllers.NewApplication(database.ProductData(database.Client,"products"),database.UserData(database.CLient,"users"))
   
	router := gin.New()
    router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart",app.AddToCart())
	router.GET("/removeitem",app.RemoveItem())
	router.GET("/checkout",app.BuyFromCart())
	router.GET("/instantbuy",app.InstantBuy())

	log.Fatal(router.Run(":"+port))

}