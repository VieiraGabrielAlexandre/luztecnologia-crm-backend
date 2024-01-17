package httprequests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"
)

func GetViaPostalCode(postalCode string) (models.Address, error) {
	baseApi := os.Getenv("VIA_CEP_URL")
	apiUrl := fmt.Sprintf(baseApi, postalCode)

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return models.Address{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Address{}, errors.New("Postal code not found")
	}

	var address models.Address
	err = json.NewDecoder(response.Body).Decode(&address)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return address, err
	}

	if address.Logradouro == "" && address.Localidade == "" && address.Cep == "" {
		return models.Address{}, errors.New("Address is entirely null")
	}

	return address, err
}
