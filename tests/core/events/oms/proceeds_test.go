package oms

// func StrToDate(dateAsText string) time.Time {
// 	layout := "2006-01-02 15:04:05"
// 	result, err := time.Parse(layout, dateAsText)

// 	if err != nil {
// 		log.Println(err)
// 		return time.Time{}
// 	}
// 	return result
// }

// func newTransactionObject(symbol string, quantity int, price float64, trade_date time.Time, broker_id float64) mapper.OMSTransaction {
// 	result := mapper.OMSTransaction{}
// 	result.Symbol = symbol
// 	result.Quantity = quantity
// 	result.Price = price
// 	result.TradeDate = trade_date
// 	result.BrokerID = broker_id
// 	return result
// }

// func newCorporateActionObject(description string, value float64, initial_date time.Time) mapper.CorporateAction {
// 	result := mapper.CorporateAction{}
// 	result.Description = description
// 	result.Value = value
// 	result.InitialDate = initial_date
// 	return result
// }

// func TestProceedsCash(t *testing.T) {

// 	// Dados chave para identificação dos ativos por usuários
// 	Customer := "WpLhDUh4"

// 	// Mock de transações
// 	OMSTransactionObjects := []mapper.OMSTransaction{}
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BBAS3", 10, 31.66, StrToDate("2021-07-06 09:40:42"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BBAS3", -10, 32.38, StrToDate("2021-07-12 10:51:08"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BBAS3", 37, 32.16, StrToDate("2021-08-09 07:17:34"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BBAS3", 5, 29.63, StrToDate("2021-08-19 07:26:58"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BBAS3", -2, 28.17, StrToDate("2021-09-21 07:18:03"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BBAS3", -10, 28.48, StrToDate("2022-01-10 07:04:48"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BBAS3", -30, 28.48, StrToDate("2022-01-10 07:04:48"), 1618))
// 	// Agrupamento das transações em um map com chave por usuário
// 	OMSTransactions := map[string][]mapper.OMSTransaction{}
// 	OMSTransactions[Customer] = OMSTransactionObjects

// 	// Mock de eventos corporativos
// 	CorporateActionObjects := []mapper.CorporateAction{}
// 	CorporateActionObjects = append(CorporateActionObjects, newCorporateActionObject("RENDIMENTO", 0.00269997702, StrToDate("2021-08-24 0:00:00")))
// 	// Agrupamento dos eventos corporativos em um map com chave por symbol
// 	CorporateActions := map[string][]mapper.CorporateAction{}
// 	CorporateActions["BBAS3"] = CorporateActionObjects

// 	// Processamento
// 	OMSProceedPersisterObject := []mapper.OMSProceeds{}
// 	OMSProceedPersisterObject = append(OMSProceedPersisterObject, oms.ApplyProceedsCorporateAction(Customer, "BBAS3", OMSTransactions, CorporateActions)...)

// 	// Teste
// 	ExpectedProceedsCount := 1
// 	if len(OMSProceedPersisterObject) != ExpectedProceedsCount {
// 		t.Errorf("TestProceeds Count: Esperado (%d), Recebido (%d)", ExpectedProceedsCount, len(OMSProceedPersisterObject))
// 	}

// 	ExpectedProceedsQuantity := 42
// 	if int(OMSProceedPersisterObject[0].Quantity) != ExpectedProceedsQuantity {
// 		t.Errorf("TestProceeds Quantity: Esperado (%d), Recebido (%d)", ExpectedProceedsQuantity, int(OMSProceedPersisterObject[0].Quantity))
// 	}

// 	ExpectedProceedsValue := 0.0
// 	if OMSProceedPersisterObject[0].Value != ExpectedProceedsValue {
// 		t.Errorf("TestProceeds Value: Esperado (%f), Recebido (%f)", ExpectedProceedsValue, OMSProceedPersisterObject[0].Value)
// 	}

// 	ExpectedProceedsAmount := 0.09
// 	if OMSProceedPersisterObject[0].Amount != ExpectedProceedsAmount {
// 		t.Errorf("TestProceeds Amount: Esperado (%f), Recebido (%f)", ExpectedProceedsAmount, OMSProceedPersisterObject[0].Amount)
// 	}

// }

// func TestProceedsBonus(t *testing.T) {

// 	// Dados chave para identificação dos ativos por usuários
// 	Customer := "2prEs4ie"

// 	// Mock de transações
// 	OMSTransactionObjects := []mapper.OMSTransaction{}
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("ITSA4", 17, 11.69, StrToDate("2021-06-11 11:42:43"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("ITSA4", 6, 11.69, StrToDate("2021-06-11 11:42:43"), 1618))
// 	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("ITSA4", 20, 11.69, StrToDate("2021-06-11 11:42:43"), 1618))
// 	// Agrupamento das transações em um map com chave por usuário
// 	OMSTransactions := map[string][]mapper.OMSTransaction{}
// 	OMSTransactions[Customer] = OMSTransactionObjects

// 	// Mock de eventos corporativos
// 	CorporateActionObjects := []mapper.CorporateAction{}
// 	CorporateActionObjects = append(CorporateActionObjects, newCorporateActionObject("BONIFICACAO", 5, StrToDate("2021-12-20 0:00:00")))
// 	// Agrupamento dos eventos corporativos em um map com chave por symbol
// 	CorporateActions := map[string][]mapper.CorporateAction{}
// 	CorporateActions["ITSA4"] = CorporateActionObjects

// 	// Processamento
// 	OMSProceedPersisterObject := []mapper.OMSProceeds{}
// 	OMSProceedPersisterObject = append(OMSProceedPersisterObject, oms.ApplyProceedsCorporateAction(Customer, "ITSA4", OMSTransactions, CorporateActions)...)

// 	// Teste
// 	ExpectedProceedsCount := 1
// 	if len(OMSProceedPersisterObject) != ExpectedProceedsCount {
// 		t.Errorf("TestProceeds Count: Esperado (%d), Recebido (%d)", ExpectedProceedsCount, len(OMSProceedPersisterObject))
// 	}

// 	ExpectedProceedsSymbol := "ITSA4"
// 	if OMSProceedPersisterObject[0].Symbol != ExpectedProceedsSymbol {
// 		t.Errorf("TestProceeds Symbol: Esperado (%s), Recebido (%s)", ExpectedProceedsSymbol, OMSProceedPersisterObject[0].Symbol)
// 	}

// 	ExpectedProceedsQuantity := 43
// 	if int(OMSProceedPersisterObject[0].Quantity) != ExpectedProceedsQuantity {
// 		t.Errorf("TestProceeds Quantity: Esperado (%d), Recebido (%d)", ExpectedProceedsQuantity, int(OMSProceedPersisterObject[0].Quantity))
// 	}

// 	ExpectedProceedsValue := 5.0
// 	if OMSProceedPersisterObject[0].Value != ExpectedProceedsValue {
// 		t.Errorf("TestProceeds Value: Esperado (%f), Recebido (%f)", ExpectedProceedsValue, OMSProceedPersisterObject[0].Value)
// 	}

// 	ExpectedProceedsAmount := 8.0
// 	if OMSProceedPersisterObject[0].Amount != ExpectedProceedsAmount {
// 		t.Errorf("TestProceeds Amount: Esperado (%f), Recebido (%f)", ExpectedProceedsAmount, OMSProceedPersisterObject[0].Amount)
// 	}

// 	ExpectedProceedsEvent := "BONIFICACAO"
// 	if OMSProceedPersisterObject[0].Event != ExpectedProceedsEvent {
// 		t.Errorf("TestProceeds Amount: Esperado (%s), Recebido (%s)", ExpectedProceedsEvent, OMSProceedPersisterObject[0].Event)
// 	}
// }
