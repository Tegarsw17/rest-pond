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
	db.AutoMigrate(&models.Fish{})

	dataUser := models.Users{}
	if db.Find(&dataUser).RecordNotFound() {
		seederAuth(db)
	}
	dataPond := models.Pond{}
	if db.Find(&dataPond).RecordNotFound() {
		seederPond(db)
	}
	dataFish := models.Fish{}
	if db.Find(&dataFish).RecordNotFound() {
		seederFish(db)
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

func seederFish(db *gorm.DB) {
	// 827ccb0eea8a706c4c34a16891f84e7b -> 12345
	date := "2022-09-09"
	parsedTimeM, err := ParseDateInput(date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	data := []models.Fish{
		{
			ID:              "1",
			Type:            "kumpay",
			Colour:          "red, white",
			Size:            "12",
			Maintenance:     "-",
			DateMaintenance: parsedTimeM,
			IdPond:          "a84e4c2a-8529-48e2-a530-6b395c2e45db",
		},
		{
			ID:              "2",
			Type:            "Shushui",
			Colour:          "Brown",
			Size:            "15",
			Maintenance:     "luka luar",
			DateMaintenance: parsedTimeM,
			IdPond:          "471ceac2-36dc-4dbc-ad8b-4422f3768250",
		},
		{
			ID:              "3",
			Type:            "ogon",
			Colour:          "Black",
			Size:            "13",
			Maintenance:     "jamur",
			DateMaintenance: parsedTimeM,
			IdPond:          "a84e4c2a-8529-48e2-a530-6b395c2e45db",
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
