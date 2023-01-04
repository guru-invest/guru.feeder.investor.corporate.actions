package domain

import (
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/utils"
	"github.com/shopspring/decimal"
)

type CorporateActionList struct {
	Map  map[string][]CorporateAction
	List []CorporateAction
}

type CorporateAction struct {
	Symbol           string
	Description      string
	Value            decimal.Decimal
	PaymentDate      time.Time
	ComDate          time.Time
	TargetTicker     string
	CalculatedFactor decimal.Decimal
	InitialDate      time.Time
}

func (h CorporateAction) IsBasic() bool {
	return utils.Contains([]string{constants.Grouping, constants.Unfolding, constants.Update}, h.Description)
}

func (h CorporateAction) IsCashProceeds() bool {
	return utils.Contains([]string{constants.InterestOnEquity, constants.Dividend, constants.Income}, h.Description)
}

func (h CorporateAction) IsBonusProceeds() bool {
	return utils.Contains([]string{constants.Bonus}, h.Description)
}
