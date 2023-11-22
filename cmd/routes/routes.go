package routes

import (
	"net/http"

	"github.com/JakeNorman007/horse_lazer/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func RoutesAll() http.Handler{
    r := gin.Default()

    r.GET("/friends", controllers.FriendsShow)

    return r
}

