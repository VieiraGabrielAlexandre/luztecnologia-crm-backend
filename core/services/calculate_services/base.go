package calculateservices

import "github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"

func Base(calc *models.Calc) {
	if calc.NumberClients <= 1999 {
		calc.Value = float64(calc.NumberClients) * 3.50
	} else {
		calc.Value = float64(calc.NumberClients) * 2.50
	}
}
