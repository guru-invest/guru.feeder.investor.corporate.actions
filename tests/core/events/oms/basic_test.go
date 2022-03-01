package oms

import (
	"testing"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

func TestBasicGrouping(t *testing.T) {
	CorporateAction := mapper.CorporateAction{}
	CorporateAction.Description = constants.Grouping
	CorporateAction.CalculatedFactor = 5

	OMSTransactionObject := mapper.OMSTransaction{}
	OMSTransactionObject.EventName = constants.Grouping
	OMSTransactionObject.Symbol = "PRIO3"
	OMSTransactionObject.Quantity = 5
	OMSTransactionObject.Price = 3.33

	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject, CorporateAction)

	ExpectedQuantity := 1.0
	if OMSTransactionObject.PostEventQuantity != ExpectedQuantity {
		t.Errorf("PostEventQuantity: Esperado (%f), Recebido (%f)", ExpectedQuantity, OMSTransactionObject.PostEventQuantity)
	}

	ExpectedPrice := 16.65
	if OMSTransactionObject.PostEventPrice != ExpectedPrice {
		t.Errorf("PostEventPrice: Esperado (%f), Recebido (%f)", ExpectedPrice, OMSTransactionObject.PostEventPrice)
	}

}

func TestBasicUnfolding(t *testing.T) {
	CorporateAction := mapper.CorporateAction{}
	CorporateAction.Description = constants.Unfolding
	CorporateAction.CalculatedFactor = 0.1996007984

	OMSTransactionObject := mapper.OMSTransaction{}
	OMSTransactionObject.EventName = constants.Unfolding
	OMSTransactionObject.Symbol = "ALUP4"
	OMSTransactionObject.Quantity = 5
	OMSTransactionObject.Price = 9.13

	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject, CorporateAction)

	ExpectedQuantity := 25.0500000004008
	if OMSTransactionObject.PostEventQuantity != ExpectedQuantity {
		t.Errorf("PostEventQuantity: Esperado (%f), Recebido (%f)", ExpectedQuantity, OMSTransactionObject.PostEventQuantity)
	}

	ExpectedPrice := 1.82
	if OMSTransactionObject.PostEventPrice != ExpectedPrice {
		t.Errorf("PostEventPrice: Esperado (%f), Recebido (%f)", ExpectedPrice, OMSTransactionObject.PostEventPrice)
	}

}

func TestBasicUpdate(t *testing.T) {
	CorporateAction := mapper.CorporateAction{}
	CorporateAction.Description = constants.Update

	OMSTransactionObject := mapper.OMSTransaction{}
	OMSTransactionObject.EventName = constants.Update
	OMSTransactionObject.Symbol = "PASS12"
	OMSTransactionObject.PostEventSymbol = "PASS5"
	OMSTransactionObject.Quantity = 10
	OMSTransactionObject.Price = 8.88
	OMSTransactionObject.EventFactor = 0

	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject, CorporateAction)

	ExpectedQuantity := 10.0
	if OMSTransactionObject.PostEventQuantity != ExpectedQuantity {
		t.Errorf("PostEventQuantity: Esperado (%f), Recebido (%f)", ExpectedQuantity, OMSTransactionObject.PostEventQuantity)
	}

	ExpectedPrice := 8.88
	if OMSTransactionObject.PostEventPrice != ExpectedPrice {
		t.Errorf("PostEventPrice: Esperado (%f), Recebido (%f)", ExpectedPrice, OMSTransactionObject.PostEventPrice)
	}

}
