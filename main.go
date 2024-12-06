package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vaisakhns/ecommerse-vs/controllers"
	"github.com/vaisakhns/ecommerse-vs/database"
	"github.com/vaisakhns/ecommerse-vs/middleware"
	"github.com/vaisakhns/ecommerse-vs/routers"
)


func main()  {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "users"))

	router := gin.New()
	
	router.Use(gin.Logger())
	routers.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeItem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.POST("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))
}

	