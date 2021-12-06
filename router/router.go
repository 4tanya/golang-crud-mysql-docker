package router

import (
    "github.com/gin-gonic/gin"

	"github.com/controllers"
)

func GetRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/books", controllers.GetAll)
    router.GET("/books/:id", controllers.GetByID)
    router.POST("/books", controllers.Save)
    router.DELETE("/books/:id", controllers.Delete)
    router.PUT("/books/:id", controllers.Update)

    return router;
}
