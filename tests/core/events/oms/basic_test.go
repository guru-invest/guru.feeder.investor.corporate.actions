package oms

// func TestBasicGrouping(t *testing.T) {
// 	CorporateAction := mapper.CorporateAction{}
// 	CorporateAction.Description = constants.Grouping
// 	CorporateAction.CalculatedFactor = 5

// 	OMSTransactionObject := mapper.OMSTransaction{}
// 	OMSTransactionObject.EventName = constants.Grouping
// 	OMSTransactionObject.Symbol = "PRIO3"
// 	OMSTransactionObject.PostEventQuantity = 5
// 	OMSTransactionObject.PostEventPrice = 3.33
// 	OMSTransactionObject.Amount = 16.65

// 	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject, CorporateAction)

// 	ExpectedQuantity := 1.0
// 	if OMSTransactionObject.PostEventQuantity != ExpectedQuantity {
// 		t.Errorf("PostEventQuantity: Esperado (%f), Recebido (%f)", ExpectedQuantity, OMSTransactionObject.PostEventQuantity)
// 	}

// 	ExpectedPrice := 16.65
// 	if OMSTransactionObject.PostEventPrice != ExpectedPrice {
// 		t.Errorf("PostEventPrice: Esperado (%f), Recebido (%f)", ExpectedPrice, OMSTransactionObject.PostEventPrice)
// 	}

// }

// func TestBasicUnfolding(t *testing.T) {
// 	CorporateAction := mapper.CorporateAction{}
// 	CorporateAction.Description = constants.Unfolding
// 	CorporateAction.CalculatedFactor = 0.1996007984

// 	OMSTransactionObject := mapper.OMSTransaction{}
// 	OMSTransactionObject.EventName = constants.Unfolding
// 	OMSTransactionObject.Symbol = "ALUP4"
// 	OMSTransactionObject.PostEventQuantity = 5
// 	OMSTransactionObject.PostEventPrice = 9.13
// 	OMSTransactionObject.Amount = 45.65

// 	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject, CorporateAction)

// 	ExpectedQuantity := 25.0500000004008
// 	if OMSTransactionObject.PostEventQuantity != ExpectedQuantity {
// 		t.Errorf("PostEventQuantity: Esperado (%f), Recebido (%f)", ExpectedQuantity, OMSTransactionObject.PostEventQuantity)
// 	}

// 	ExpectedPrice := 1.82
// 	if OMSTransactionObject.PostEventPrice != ExpectedPrice {
// 		t.Errorf("PostEventPrice: Esperado (%f), Recebido (%f)", ExpectedPrice, OMSTransactionObject.PostEventPrice)
// 	}

// }

// func TestBasicUpdate(t *testing.T) {
// 	CorporateAction := mapper.CorporateAction{}
// 	CorporateAction.Description = constants.Update

// 	OMSTransactionObject := mapper.OMSTransaction{}
// 	OMSTransactionObject.EventName = constants.Update
// 	OMSTransactionObject.Symbol = "PASS12"
// 	OMSTransactionObject.PostEventSymbol = "PASS5"
// 	OMSTransactionObject.PostEventQuantity = 10
// 	OMSTransactionObject.PostEventPrice = 8.88
// 	OMSTransactionObject.EventFactor = 0
// 	OMSTransactionObject.Amount = 88.8

// 	OMSTransactionObject = oms.ApplyBasicCorporateAction(OMSTransactionObject, CorporateAction)

// 	ExpectedQuantity := 10.0
// 	if OMSTransactionObject.PostEventQuantity != ExpectedQuantity {
// 		t.Errorf("PostEventQuantity: Esperado (%f), Recebido (%f)", ExpectedQuantity, OMSTransactionObject.PostEventQuantity)
// 	}

// 	ExpectedPrice := 8.88
// 	if OMSTransactionObject.PostEventPrice != ExpectedPrice {
// 		t.Errorf("PostEventPrice: Esperado (%f), Recebido (%f)", ExpectedPrice, OMSTransactionObject.PostEventPrice)
// 	}

// }
