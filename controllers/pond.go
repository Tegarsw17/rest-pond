package controllers

import (
	"fmt"
	"log"
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

	// h.DB.Find(&ponds, "id_user = ?", "9daaacda-0f24-4f67-aba3-e779a2ddcc82")

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	str := fmt.Sprintf("%v", userID)

	h.DB.Find(&ponds, "id_user = ?", str)

	if len(ponds) == 0 {
		c.JSON(http.StatusNotFound, utils.ResponsJsonStruct{
			Error:   true,
			Message: "Data not found",
			Data:    ponds,
		})
		return
	}

	var pondResponses []models.ShowPond

	for _, pond := range ponds {
		pondResponses = append(pondResponses, models.ShowPond{
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
		Message: "All Pond",
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
		Data:    payload.ID,
	})
}

func (h *Handler) GetDetailPond(c *gin.Context) {
	var pond models.Pond
	id := c.Param("id")
	log.Println(id)
	h.DB.First(&pond, "id = ?", id)
	if pond.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Pond not found",
			Data:    "",
		})
		return
	}
	data := models.ShowPond{
		ID:              pond.ID,
		Name:            pond.Name,
		Dimension:       pond.Dimension,
		Condition:       pond.Condition,
		Maintenance:     pond.Maintenance,
		DateMaintenance: pond.DateMaintenance,
		DateFeeding:     pond.DateFeeding,
		TotalFish:       pond.TotalFish,
	}
	// fmt.Print(pond)
	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "Pond found",
		Data:    data,
	})
}

func (h *Handler) DeletePond(c *gin.Context) {
	var pond models.Pond
	id := c.Param("id")
	h.DB.First(&pond, "id = ?", id)
	if pond.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Pond not found",
			Data:    "",
		})
		return
	}
	h.DB.Where("id = ?", id).Delete(&pond)
	fmt.Print(pond)
	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: "Pond succes delete",
		Data:    id,
	})
}

func (h *Handler) UpdatePond(c *gin.Context) {
	var payload models.CreatePond
	var pond models.Pond

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	str := fmt.Sprintf("%v", userID)

	err := c.Bind(&payload)
	// fmt.Print(payload)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	id := c.Param("id")
	h.DB.First(&pond, "id = ?", id)
	if pond.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Pond not found",
			Data:    "",
		})
		return
	}

	parsedTimeM, err := utils.ParseDateInput(payload.DateMaintenance)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	parsedTimeF, err := utils.ParseDateInput(payload.DateFeeding)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	pond.Name = payload.Name
	pond.Dimension = payload.Dimension
	pond.Condition = payload.Condition
	pond.Maintenance = payload.Maintenance
	pond.DateMaintenance = parsedTimeM
	pond.DateFeeding = parsedTimeF
	pond.TotalFish = payload.TotalFish
	pond.IdUser = str

	h.DB.Save(&pond)
	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: "pond succes update",
		Data:    pond.ID,
	})
}
