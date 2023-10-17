package routers

import (
	"pond-manage/controllers"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PondRouter(r *gin.Engine, db *gorm.DB) {
	handler := controllers.New(db)
	// pond := r.Group("/api")
	// pond.Use(utils.JWTMiddleware())
	// {
	// }
	r.GET("/pond", utils.JWTMiddleware(), handler.GetAllPond)
	r.POST("/pond", utils.JWTMiddleware(), handler.CreatePond)
}
