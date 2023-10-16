package controllers

import (
	"net/http"
	"pond-manage/models"
	"pond-manage/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPond(c *gin.Context) {
	var ponds []models.Pond
	h.DB.Find(&ponds, "id_user = ?", "9daaacda-0f24-4f67-aba3-e779a2ddcc82")

	// response := PondResponse{
	// 	Ponds: ponds,
	// }

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
