package domain

import "time"

type CorporateActions struct {
	Symbol          string
	CorporateAction []CorporateActionDetail
}

type CorporateActionDetail struct {
	Symbol           string
	Description      string
	Value            float64
	PaymentDate      time.Time
	ComDate          time.Time
	TargetTicker     string
	CalculatedFactor float64
	InitialDate      time.Time
}

func (t CorporateActions) Create(symbol string, corporateAction []CorporateActionDetail) CorporateActions {

	return CorporateActions{
		Symbol:          symbol,
		CorporateAction: corporateAction,
	}
}

func (t CorporateActionDetail) Create(symbol string, description string, value float64, paymentDate time.Time, comDate time.Time, targetTicker string, calculatedFactor float64, initialDate time.Time) CorporateActionDetail {

	return CorporateActionDetail{
		Symbol:           symbol,
		Description:      description,
		Value:            value,
		PaymentDate:      paymentDate,
		ComDate:          comDate,
		TargetTicker:     targetTicker,
		CalculatedFactor: calculatedFactor,
		InitialDate:      initialDate,
	}
}
