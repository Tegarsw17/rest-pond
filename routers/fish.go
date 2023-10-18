package routers

import (
	"pond-manage/controllers"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FishRouter(r *gin.Engine, db *gorm.DB) {
	handler := controllers.New(db)

	r.GET("/pond/:id/fish", utils.JWTMiddleware(), handler.GetAllFish)

	r.POST("/pond/:id/fish", utils.JWTMiddleware(), handler.CreateFish)
	r.GET("/pond/:id/fish/:id_fish", utils.JWTMiddleware(), handler.GetDetailFish)
	r.PUT("/pond/:id/fish/:id_fish", utils.JWTMiddleware(), handler.UpdateFish)
	r.DELETE("/pond/:id/fish/:id_fish", utils.JWTMiddleware(), handler.DeleteFish)
}
