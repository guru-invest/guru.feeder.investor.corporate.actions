package mapper

func (Customer) TableName() string {
	return "wallet.cei_items_status"
}

type Customer struct {
	CustomerCode string `gorm:"column:customer_code"`
}
