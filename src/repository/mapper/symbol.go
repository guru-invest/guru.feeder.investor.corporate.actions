package mapper

func (Symbol) TableName() string {
	return "wallet.oms_transactions"
}

type Symbol struct {
	Name string `gorm:"column:symbol"`
}
