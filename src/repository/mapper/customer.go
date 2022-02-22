package mapper

func (Customer) TableName() string {
	return "wallet.oms_transactions"
}

type Customer struct {
	CustomerCode string `gorm:"column:customer_code"`
}
