package controllers

import (
	"fmt"
	"net/http"
	"pond-manage/models"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func (h *Handler) GetAllFish(c *gin.Context) {
	var fishs []models.Fish
	id := c.Param("id")
	// h.DB.Find(&fishs, "id_pond = ?", "a84e4c2a-8529-48e2-a530-6b395c2e45db")
	h.DB.Find(&fishs, "id_pond = ?", id)

	if len(fishs) == 0 {
		c.JSON(http.StatusNotFound, utils.ResponsJsonStruct{
			Error:   true,
			Message: "Data not found",
			Data:    fishs,
		})
		return
	}

	var fishRespond []models.ShowFish

	for _, fish := range fishs {
		fishRespond = append(fishRespond, models.ShowFish{
			ID:              fish.ID,
			Type:            fish.Type,
			Colour:          fish.Colour,
			Size:            fish.Size,
			Maintenance:     fish.Maintenance,
			DateMaintenance: fish.DateMaintenance,
			IdPond:          fish.IdPond,
		})

	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "All Fish",
		Data:    fishRespond,
	})
}

func (h *Handler) CreateFish(c *gin.Context) {
	var payload models.CreatedFish
	var fish models.Fish

	err := c.Bind(&payload)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	id_pond := c.Param("id")
	id := uuid.NewV4()
	parsedTime, err := utils.ParseDateInput(payload.DateMaintenance)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fish.ID = id.String()
	fish.Type = payload.Type
	fish.Colour = payload.Colour
	fish.Size = payload.Size
	fish.Maintenance = payload.Maintenance
	fish.DateMaintenance = parsedTime
	fish.IdPond = id_pond

	h.DB.Save(&fish)
	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: "fish created succes",
		Data:    fish.ID,
	})
}

func (h *Handler) GetDetailFish(c *gin.Context) {
	var fish models.Fish

	id_pond := c.Param("id")
	id_fish := c.Param("id_fish")

	h.DB.Where("id = ? AND id_pond = ?", id_fish, id_pond).First(&fish)
	if fish.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "fish not found",
			Data:    "",
		})
		return
	}
	data := models.ShowFish{
		ID:              fish.ID,
		Type:            fish.Type,
		Colour:          fish.Colour,
		Size:            fish.Size,
		Maintenance:     fish.Maintenance,
		DateMaintenance: fish.DateMaintenance,
		IdPond:          fish.IdPond,
	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "fish found",
		Data:    data,
	})
}

func (h *Handler) UpdateFish(c *gin.Context) {
	var payload models.CreatedFish
	var fish models.Fish

	err := c.Bind(&payload)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	id_pond := c.Param("id")
	id_fish := c.Param("id_fish")

	h.DB.Where("id = ? AND id_pond = ?", id_fish, id_pond).First(&fish)
	if fish.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "fish not found",
			Data:    "",
		})
		return
	}
	parsedTime, err := utils.ParseDateInput(payload.DateMaintenance)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	fish.Type = payload.Type
	fish.Colour = payload.Colour
	fish.Size = payload.Size
	fish.Maintenance = payload.Maintenance
	fish.DateMaintenance = parsedTime
	fish.IdPond = id_pond

	h.DB.Save(&fish)
	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: "pond succes update",
		Data:    fish.ID,
	})

}

func (h *Handler) DeleteFish(c *gin.Context) {
	var fish models.Fish
	id_pond := c.Param("id")
	id_fish := c.Param("id_fish")

	h.DB.Where("id = ? AND id_pond = ?", id_fish, id_pond).First(&fish)
	if fish.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "fish not found",
			Data:    "",
		})
		return
	}

	h.DB.Where("id = ? AND id_pond = ?", id_fish, id_pond).Delete(&fish)
	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: "fish succes delete",
		Data:    id_fish,
	})
}
