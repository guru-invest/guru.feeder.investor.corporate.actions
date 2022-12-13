package domain

import "time"

type CorporateActionList struct {
	Map  map[string][]CorporateAction
	List []CorporateAction
}

type CorporateAction struct {
	Symbol           string
	Description      string
	Value            float64
	PaymentDate      time.Time
	ComDate          time.Time
	TargetTicker     string
	CalculatedFactor float64
	InitialDate      time.Time
}
