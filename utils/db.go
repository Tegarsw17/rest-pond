package utils

import (
	"fmt"
	"log"
	"os"
	"pond-manage/models"
	"time"

	_ "database/sql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	Migrate(db)
	return db, nil

}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Pond{})

	dataUser := models.Users{}
	if db.Find(&dataUser).RecordNotFound() {
		seederAuth(db)
	}
	dataPond := models.Pond{}
	if db.Find(&dataPond).RecordNotFound() {
		seederPond(db)
	}
}

func seederAuth(db *gorm.DB) {
	// 827ccb0eea8a706c4c34a16891f84e7b -> 12345
	data := []models.Users{
		{
			ID:       "1",
			Username: "tegar",
			Password: "827ccb0eea8a706c4c34a16891f84e7b",
			Email:    "tegar@mail.com",
		},
		{
			ID:       "2",
			Username: "satriya",
			Password: "827ccb0eea8a706c4c34a16891f84e7b",
			Email:    "satriya@mail.com",
		},
	}

	for _, v := range data {
		db.Create(&v)
	}
}

func seederPond(db *gorm.DB) {
	dateMaintenanceStr := "2023-10-20 12:30:00"
	dateMaintenance, err := time.Parse("2006-01-02 15:04:05", dateMaintenanceStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	data := []models.Pond{
		{
			ID:              "1",
			Name:            "Depan rumah",
			Dimension:       "1.2, 2.5, 3",
			Condition:       "bad",
			Maintenance:     "change spons",
			DateMaintenance: dateMaintenance,
			DateFeeding:     dateMaintenance,
			TotalFish:       5,
			IdUser:          "9daaacda-0f24-4f67-aba3-e779a2ddcc82",
		},
		{
			ID:              "2",
			Name:            "taman",
			Dimension:       "1.2, 2.5, 3",
			Condition:       "good",
			Maintenance:     "cleaning",
			DateMaintenance: dateMaintenance,
			DateFeeding:     dateMaintenance,
			TotalFish:       15,
			IdUser:          "c4356847-291f-4b0b-a845-31910e019bfe",
		},
		{
			ID:              "3",
			Name:            "rooftop",
			Dimension:       "1.2, 2.5, 3",
			Condition:       "good",
			Maintenance:     "cleaning",
			DateMaintenance: dateMaintenance,
			DateFeeding:     dateMaintenance,
			TotalFish:       24,
			IdUser:          "9daaacda-0f24-4f67-aba3-e779a2ddcc82",
		},
	}

	for _, v := range data {
		db.Create(&v)
	}
}
