package main

import (
	"pond-manage/routers"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := utils.SetupDB()
	if err != nil {
		// Handle the error, e.g., log it and exit
		panic(err)
	}
	routers.AuthRouter(r, db)
	routers.PondRouter(r, db)
	routers.FishRouter(r, db)
	r.Run("localhost:8080")
}
