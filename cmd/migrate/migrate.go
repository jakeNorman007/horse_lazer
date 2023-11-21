package main

import (
    "github.com/JakeNorman007/horse_lazer/cmd/models"
    "github.com/JakeNorman007/horse_lazer/cmd/initializers"
)

func init(){
    initializers.LoadEnvVar()
    initializers.ConnectToDB()
}

func main(){
    initializers.DB.AutoMigrate(&models.Friend{})
}
