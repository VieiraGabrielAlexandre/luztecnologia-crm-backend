package clientsrepository

import (
	"errors"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/config"
	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"
	"gorm.io/gorm"
)

func Save(client *models.Client) {
	if err := config.DB.Model(&client).Where("cnpj = ?", client.Cnpj).Updates(client).Error; err != nil {
		// always handle error like this, cause errors maybe happened when connection failed or something.
		// record not found...
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.DB.Create(&client) // create new record from newUser
		}
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
