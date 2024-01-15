package models

type Calc struct {
	NumberClients int     `json:"number_clients" type:"int"`
	Value         float64 `json:"value" type:"float64"`
}
