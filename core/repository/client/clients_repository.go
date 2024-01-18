package clientsrepository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/config"
	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"
)

func Save(client *models.Client) (string, error) {
	var existingClient models.Client
	if err := config.DB.Where("id = ? OR cnpj = ?", client.ID, client.Cnpj).First(&existingClient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.DB.Create(&client)
			return "Client created", nil

		} else {
			fmt.Println("Error:", err)
			return "", err
		}
	} else {
		config.DB.Model(&existingClient).Updates(client)
		return "Client updated", nil
	}

}

func GetAll() []models.Client {
	var clients []models.Client
	config.DB.Find(&clients)
	return clients
}

func Get(id int) (models.Client, error) {
	var client models.Client
	result := config.DB.Where("id = ?", id).First(&client)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Handle the case where the record is not found
			return models.Client{}, errors.New("Client not found")
		}
		// Handle other errors if needed
		return models.Client{}, result.Error
	}

	return client, nil
}
func Count() int64 {
	var count int64
	config.DB.Model(&models.Client{}).Count(&count)
	return count
}
