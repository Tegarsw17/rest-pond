package controllers

import (
	"log"
	"net/http"
	"os"
	"pond-manage/models"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func (h *Handler) LoginTask(c *gin.Context) {
	var credential models.Login
	var user models.Users
	err := c.Bind(&credential)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	if !h.isUsernameAvailabel(credential.Username) {
		c.JSON(http.StatusBadRequest, utils.ResponsJson{
			Error:   true,
			Message: "username is not found",
			Data:    "",
		})
		return
	}

	h.DB.First(&user, "username=?", credential.Username)

	if !h.isUsernameAvailabel(credential.Username) || !CheckPassword(credential.Password, user.Password) {
		c.JSON(http.StatusBadRequest, utils.ResponsJson{
			Error:   true,
			Message: "username or password is invalid",
			Data:    "",
		})
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error load env")
		}
		h.DB.First(&user, "username=?", credential.Username)
		secret := os.Getenv("JWT_SECRET")
		token, err := utils.GenerateToken(user.ID, secret)

		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ResponsJson{
				Error:   true,
				Message: "login failed",
				Data:    "",
			})
		}

		c.JSON(http.StatusOK, utils.ResponsJson{
			Error:   false,
			Message: "login succes",
			Data:    token,
		})

	}

}

func CheckPassword(password string, hashedPassword string) bool {
	return hashedPassword == hashPass(password)
}
