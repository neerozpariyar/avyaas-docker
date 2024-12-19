package main

import (
	"avyaas/internal/config"
	"avyaas/internal/domain/models"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	config.ConfigureViper()
	db := config.InitDB(viper.GetBool("verbose"), viper.GetBool("logger"))

	packageTypes := []models.PackageType{
		{
			Timestamp: models.Timestamp{
				ID: 1,
			},
			Title:       "Combo (Course + Test Series + Live Group)",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 2,
			},
			Title:       "Combo (Course + Test Series)",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 3,
			},
			Title:       "Combo (Course + Live Group)",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 4,
			},
			Title:       "Combo (Test Series + Live Group)",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 5,
			},
			Title:       "Course",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 6,
			},
			Title:       "Test Series",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 7,
			},
			Title:       "Live Group",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 8,
			},
			Title:       "Single Test",
			Description: "",
		},
		{
			Timestamp: models.Timestamp{
				ID: 9,
			},
			Title:       "Single Live",
			Description: "",
		},
	}

	for _, packageType := range packageTypes {
		err := db.Where("name = ?", packageType.DeletedAt.Time).First(&models.PackageType{}).Error
		if err != nil {
			db.Create(&packageType)
			fmt.Printf("[+][+] Package type created for: %s [+][+]\n", packageType.Title)
		}
	}
	fmt.Println("[+][+] All Package type created [+][+]")
}
