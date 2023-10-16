package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"pond-manage/models"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) RegisterTask(c *gin.Context) {
	var payload models.Users
	err := c.Bind(&payload)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if payload.Username == "" || payload.Email == "" {
		c.JSON(http.StatusBadRequest, utils.ResponsJson{
			Error:   true,
			Message: "Input is invalid",
			Data:    "",
		})
		return
	}

	isDuplicateUsername := h.isUsernameAvailabel(payload.Username)
	isDuplicateEmail := h.isEmailDuplicate(payload.Email)

	if isDuplicateUsername {
		c.JSON(http.StatusBadRequest, utils.ResponsJson{
			Error:   true,
			Message: "username already exist",
			Data:    "",
		})
		return
	}

	if isDuplicateEmail {
		c.JSON(http.StatusBadRequest, utils.ResponsJson{
			Error:   true,
			Message: "email already exist",
			Data:    "",
		})
		return
	}

	id := uuid.NewV4()
	payload.ID = id.String()
	payload.Password = hashPass(payload.Password)
	h.DB.Create(&payload)
	log.Print(payload.Username)
	dataMap := map[string]interface{}{
		"username": payload.Username,
	}
	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "login succes",
		Data:    dataMap,
	})

}

func (h *Handler) isUsernameAvailabel(username string) bool {
	var user models.Users
	h.DB.Where("username = ?", username).First(&user)
	return user.ID != ""
}

func (h *Handler) isEmailDuplicate(email string) bool {
	var user models.Users
	h.DB.Where("email = ?", email).First(&user)
	return user.ID != ""
}

func hashPass(password string) string {
	inputBytes := []byte(password)
	md5Hash := md5.Sum(inputBytes)
	md5HashString := hex.EncodeToString(md5Hash[:])
	return md5HashString
}
