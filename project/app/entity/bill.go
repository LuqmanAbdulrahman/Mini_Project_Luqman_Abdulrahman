package entity

type Bill struct {
	ID           uint     `json:"id" gorm:"primaryKey"`
	CustomerID   uint     `json:"customer_id"`
	Customer     Customer `json:"customer" gorm:"foreignKey:CustomerID"`
	Period       string   `json:"period"`
	CubicMeter   float64  `json:"cubic_meter"`
	TotalPrice   float64  `json:"total_price"`
	PaymentState bool     `json:"payment_state"`
}
