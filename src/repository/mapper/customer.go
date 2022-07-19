package mapper

type Customer struct {
	CustomerCode string `gorm:"column:customer_code"`
	CreatedAT    string `gorm:"column:created_at"`
}
