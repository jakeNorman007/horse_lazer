package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//"github.com/JakeNorman007/horse_lazer/cmd/controllers"
//"github.com/JakeNorman007/horse_lazer/cmd/initializers"
//"github.com/gin-gonic/gin"

//func init(){
//initializers.LoadEnvVar()
//initializers.ConnectToDB()
//}

type Friend struct{
    Name    string
    Email   string
    Phone   string
}

func main() {
    fmt.Println("Server connected.")

    handlerOne := func(w http.ResponseWriter, r *http.Request){
        tmpl := template.Must(template.ParseFiles("src/views/index.html"))
        friends := map[string][]Friend{
            "friends": {
                {Name: "Jake Norman", Email: "jake@example.com", Phone: "555-555-5555"},
                {Name: "Kazmir Dinse", Email: "kazmir@example.com", Phone: "444-444-4444"},
                {Name: "Elliott Norman", Email: "elliott@example.com", Phone: "333-333-3333"},
            },
        }
        tmpl.Execute(w, friends)
    }

    http.HandleFunc("/", handlerOne)

    log.Fatal(http.ListenAndServe(":42069", nil))
    //r := gin.Default()

    // routes, will probably put these in their own file
    //r.GET("/friends", controllers.FriendsShow)
    //r.GET("/friend/:id", controllers.FriendShowById)
    //r.POST("/friend", controllers.FriendCreate)
    //r.PUT("/friend/:id", controllers.FriendUpdate)
    //r.DELETE("/friend/:id", controllers.FriendDelete)

    //r.Run() //localhost:42069
}
