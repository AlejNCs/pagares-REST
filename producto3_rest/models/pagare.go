package models

type Pagare struct {
	ID     int     `json:"id"`
	Monto  float64 `json:"monto"`
	Meses  int     `json:"meses"`
	Status string  `json:"status"`
}
