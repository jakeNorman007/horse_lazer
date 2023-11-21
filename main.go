package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/JakeNorman007/horse_lazer/cmd/controllers"
	"github.com/JakeNorman007/horse_lazer/cmd/initializers"
)

func init(){
    initializers.LoadEnvVar()
    initializers.ConnectToDB()
}

func main() {
    r := gin.Default()
    
    r.LoadHTMLGlob("src/*")

    r.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "index.html", nil)
    })
    r.Run(":42069")

    // going to use static data for this project to keep it extra simple
    // so I can just mess with some HTMX for fun
    // creating database was more for prcatice for an upcoming project
    //
    // r.GET("/", controllers.FriendsShow)
    // r.GET("/:id", controllers.FriendShowById)
    // r.POST("/", controllers.FriendCreate)
    // r.PUT("/:id", controllers.FriendUpdate)
    // r.DELETE("/:id", controllers.FriendDelete)

    // r.Run() //localhost:42069
}
