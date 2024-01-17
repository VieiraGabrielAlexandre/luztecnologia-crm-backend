package viaceprepository

import (
	httprequests "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/https_requests"
	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"
)

func ApiGetByPostalCode(postalCode string) (models.Address, error) {
	address, err := httprequests.GetViaPostalCode(postalCode)

	if err != nil {
		return models.Address{}, err
	}

	return address, nil
}
