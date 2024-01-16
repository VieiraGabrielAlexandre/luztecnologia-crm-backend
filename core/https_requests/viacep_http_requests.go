package httprequests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func GetViaPostalCode(string postalCode) {
	apiUrl := fmt.Sprintf(os.GetEnv("VIA_CEP_URL") . postalCode)	

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	var address Address
	err = json.NewDecoder(response.Body).Decode(&address)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	fmt.Printf("CEP: %s\n", address.Cep)
	fmt.Printf("Logradouro: %s\n", address.Logradouro)
	fmt.Printf("Bairro: %s\n", address.Bairro)
	fmt.Printf("Localidade: %s\n", address.Localidade)
	fmt.Printf("UF: %s\n", address.UF)
}