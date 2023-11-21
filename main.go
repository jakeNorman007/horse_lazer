package main

import (
	"github.com/gin-gonic/gin"
	"github.com/JakeNorman007/horse_lazer/cmd/controllers"
	"github.com/JakeNorman007/horse_lazer/cmd/initializers"
)

func init(){
    initializers.LoadEnvVar()
    initializers.ConnectToDB()
}

func main() {
    r := gin.Default()

    r.GET("/", controllers.FriendsShow)
    r.GET("/:id", controllers.FriendShowById)
    r.POST("/", controllers.FriendCreate)
    r.PUT("/:id", controllers.FriendUpdate)
    r.DELETE("/:id", controllers.FriendDelete)

    r.Run() //localhost:42069
}
