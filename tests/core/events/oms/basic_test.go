package oms

import (
	"testing"

	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
)

func TestBasicGrouping(t *testing.T) {

	OMSTransactionObject := mapper.OMSTransaction{}
	OMSTransactionObject.EventName = singleton.New().Grouping
	OMSTransactionObject.Symbol = "PRIO3"
	OMSTransactionObject.Quantity = 5
	OMSTransactionObject.Price = 3.33
	OMSTransactionObject.EventFactor = 5

	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject)

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

	OMSTransactionObject := mapper.OMSTransaction{}
	OMSTransactionObject.EventName = singleton.New().Unfolding
	OMSTransactionObject.Symbol = "ALUP4"
	OMSTransactionObject.Quantity = 5
	OMSTransactionObject.Price = 9.13
	OMSTransactionObject.EventFactor = 0.1996007984

	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject)

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

	OMSTransactionObject := mapper.OMSTransaction{}
	OMSTransactionObject.EventName = singleton.New().Update
	OMSTransactionObject.Symbol = "PASS12"
	OMSTransactionObject.PostEventSymbol = "PASS5"
	OMSTransactionObject.Quantity = 10
	OMSTransactionObject.Price = 8.88
	OMSTransactionObject.EventFactor = 0

	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject)

	ExpectedQuantity := 10.0
	if OMSTransactionObject.PostEventQuantity != ExpectedQuantity {
		t.Errorf("PostEventQuantity: Esperado (%f), Recebido (%f)", ExpectedQuantity, OMSTransactionObject.PostEventQuantity)
	}

	ExpectedPrice := 8.88
	if OMSTransactionObject.PostEventPrice != ExpectedPrice {
		t.Errorf("PostEventPrice: Esperado (%f), Recebido (%f)", ExpectedPrice, OMSTransactionObject.PostEventPrice)
	}

}
