package controllers

import(
    "github.com/gin-gonic/gin"
    "github.com/JakeNorman007/horse_lazer/cmd/models"
    "github.com/JakeNorman007/horse_lazer/cmd/initializers"
)

func FriendCreate(c *gin.Context){
    // get the data off of the request body
    var body struct{
        Name    string
        Email   string
        Phone   int
    }

    //converts the body of the "friend" to JSON
    c.Bind(&body)

    friend := models.Friend{Name: body.Name, Email: body.Email, Phone: body.Phone}

    result := initializers.DB.Create(&friend)

    if result != nil{
        c.Status(200)
        return
    }

    // returns the friend you have created
    c.JSON(200, gin.H{
        "friend": friend,
    })
}

func FriendsShow(c *gin.Context){
    
    var friends []models.Friend
    initializers.DB.Find(&friends)

    c.JSON(200, friends)
}

func FriendShowById(c *gin.Context){
    
    id := c.Param("id")

    var friend models.Friend
    initializers.DB.Find(&friend, id)

    c.JSON(200, gin.H{
        "friend": friend,
    })
}

func FriendUpdate(c *gin.Context){
    
    id := c.Param("id")

    var body struct{
        Name    string
        Email   string
        Phone   int
    }

    c.Bind(&body)

    var friend models.Friend
    initializers.DB.First(&friend, id)

    initializers.DB.Model(&friend).Updates(models.Friend{Name: body.Name, Email: body.Email, Phone: body.Phone,})

    c.JSON(200, gin.H{
        "friend": friend,
    })
}

func FriendDelete(c *gin.Context){
    
    id := c.Param("id")

    initializers.DB.Delete(&models.Friend{}, id)

    c.Status(200)
}
