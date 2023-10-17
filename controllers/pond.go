package controllers

import (
	"fmt"
	"net/http"
	"pond-manage/models"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	uuid "github.com/satori/go.uuid"
)

var validate = validator.New()
var english = en.New()
var uni = ut.New(english, english)

func (h *Handler) GetAllPond(c *gin.Context) {

	var ponds []models.Pond
	h.DB.Find(&ponds, "id_user = ?", "9daaacda-0f24-4f67-aba3-e779a2ddcc82")

	var pondResponses []models.CreatePond

	for _, pond := range ponds {
		pondResponses = append(pondResponses, models.CreatePond{
			ID:              pond.ID,
			Name:            pond.Name,
			Dimension:       pond.Dimension,
			Condition:       pond.Condition,
			Maintenance:     pond.Maintenance,
			DateMaintenance: pond.DateMaintenance,
			DateFeeding:     pond.DateFeeding,
			TotalFish:       pond.TotalFish,
		})
	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "ini adalah get pond",
		Data:    pondResponses,
	})
}

func (h *Handler) CreatePond(c *gin.Context) {
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	var payload models.Pond

	err := c.Bind(&payload)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	str := fmt.Sprintf("%v", userID)
	id := uuid.NewV4()
	payload.ID = id.String()
	payload.IdUser = str

	if err := validate.Struct(payload); err != nil {
		errs := utils.TranslateError(err, trans)
		c.JSON(http.StatusOK, utils.ResponsJsonStruct{
			Error:   true,
			Message: "created pond failed",
			Data:    errs,
		})
		return
	}

	if err := h.DB.Create(&payload).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "create"})
		return
	}

	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: "pond created succes",
		Data:    payload.Name,
	})
}
