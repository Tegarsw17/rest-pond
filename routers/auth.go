package routers

import (
	"pond-manage/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AuthRouter(r *gin.Engine, db *gorm.DB) {
	handler := controllers.New(db)
	r.POST("/login", handler.LoginTask)
	r.POST("/register", handler.RegisterTask)
}
