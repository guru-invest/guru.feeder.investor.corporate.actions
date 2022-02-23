package oms

import (
	"log"
	"testing"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

func StrToDate(DateAsText string) time.Time {
	layout := "2006-01-02 15:04:05"
	result, err := time.Parse(layout, DateAsText)

	if err != nil {
		log.Println(err)
		return time.Time{}
	}
	return result
}

func TestProceeds(t *testing.T) {

	// Dados chave para identificação dos ativos por usuários
	Customer := "fzVzgo8b"
	Symbol_1 := "BIDI4"

	// Mock de transações
	OMSTransactionObjects := []mapper.OMSTransaction{}

	OMSTransactionObject1 := mapper.OMSTransaction{}
	OMSTransactionObject1.Symbol = Symbol_1
	OMSTransactionObject1.Quantity = 3
	OMSTransactionObject1.Price = 27.41
	OMSTransactionObject1.TradeDate = StrToDate("2021-07-20 11:40:26")
	OMSTransactionObjects = append(OMSTransactionObjects, OMSTransactionObject1)

	OMSTransactionObject2 := mapper.OMSTransaction{}
	OMSTransactionObject2.Symbol = Symbol_1
	OMSTransactionObject2.Quantity = 5
	OMSTransactionObject2.Price = 26.68
	OMSTransactionObject2.TradeDate = StrToDate("2021-07-26 10:29:32")
	OMSTransactionObjects = append(OMSTransactionObjects, OMSTransactionObject2)

	OMSTransactionObject3 := mapper.OMSTransaction{}
	OMSTransactionObject3.Symbol = Symbol_1
	OMSTransactionObject3.Quantity = 29
	OMSTransactionObject3.Price = 26.68
	OMSTransactionObject3.TradeDate = StrToDate("2021-07-26 10:29:32")
	OMSTransactionObjects = append(OMSTransactionObjects, OMSTransactionObject3)

	OMSTransactionObject4 := mapper.OMSTransaction{}
	OMSTransactionObject4.Symbol = Symbol_1
	OMSTransactionObject4.Quantity = 1
	OMSTransactionObject4.Price = 25.69
	OMSTransactionObject4.TradeDate = StrToDate("2021-07-28 8:11:05")
	OMSTransactionObjects = append(OMSTransactionObjects, OMSTransactionObject4)

	OMSTransactionObject5 := mapper.OMSTransaction{}
	OMSTransactionObject5.Symbol = Symbol_1
	OMSTransactionObject5.Quantity = 1
	OMSTransactionObject5.Price = 25.67
	OMSTransactionObject5.TradeDate = StrToDate("2021-07-28 8:13:40")
	OMSTransactionObjects = append(OMSTransactionObjects, OMSTransactionObject5)

	OMSTransactionObject6 := mapper.OMSTransaction{}
	OMSTransactionObject6.Symbol = Symbol_1
	OMSTransactionObject6.Quantity = 20
	OMSTransactionObject6.Price = 9.47
	OMSTransactionObject6.TradeDate = StrToDate("2022-01-03 11:34:38")
	OMSTransactionObjects = append(OMSTransactionObjects, OMSTransactionObject6)

	// Agrupamento das transações em um map com chave por usuário
	OMSTransactions := map[string][]mapper.OMSTransaction{}
	OMSTransactions[Customer] = OMSTransactionObjects

	// Mock de eventos corporativos
	CorporateActionObjects := []mapper.CorporateAction{}

	CorporateActionObject1 := mapper.CorporateAction{}
	CorporateActionObject1.ComDate = StrToDate("2021-07-22 0:00:00")
	CorporateActionObject1.Value = 0.012084364
	CorporateActionObject1.Description = "JRS CAP PROPRIO"
	CorporateActionObjects = append(CorporateActionObjects, CorporateActionObject1)

	// Agrupamento dos eventos corporativos em um map com chave por symbol
	CorporateActions := map[string][]mapper.CorporateAction{}
	CorporateActions[Symbol_1] = CorporateActionObjects

	OMSProceedPersisterObject := []mapper.OMSProceeds{}
	OMSProceedPersisterObject = append(OMSProceedPersisterObject, oms.ApplyCashProceedsCorporateAction(Customer, Symbol_1, OMSTransactions, CorporateActions)...)

	ExpectedProceedsCount := 1
	if len(OMSProceedPersisterObject) != ExpectedProceedsCount {
		t.Errorf("TestProceeds Count: Esperado (%d), Recebido (%d)", ExpectedProceedsCount, len(OMSProceedPersisterObject))
	}

	ExpectedProceedsSymbol := "BIDI4"
	if OMSProceedPersisterObject[0].Symbol != ExpectedProceedsSymbol {
		t.Errorf("TestProceeds Symbol: Esperado (%s), Recebido (%s)", ExpectedProceedsSymbol, OMSProceedPersisterObject[0].Symbol)
	}

	ExpectedProceedsQuantity := 3
	if int(OMSProceedPersisterObject[0].Quantity) != ExpectedProceedsQuantity {
		t.Errorf("TestProceeds Quantity: Esperado (%d), Recebido (%d)", ExpectedProceedsQuantity, int(OMSProceedPersisterObject[0].Quantity))
	}

	ExpectedProceedsValue := 0.012084364
	if OMSProceedPersisterObject[0].Value != ExpectedProceedsValue {
		t.Errorf("TestProceeds Value: Esperado (%f), Recebido (%f)", ExpectedProceedsValue, OMSProceedPersisterObject[0].Value)
	}

	ExpectedProceedsAmount := 0.036253092
	if OMSProceedPersisterObject[0].Amount != ExpectedProceedsAmount {
		t.Errorf("TestProceeds Amount: Esperado (%f), Recebido (%f)", ExpectedProceedsAmount, OMSProceedPersisterObject[0].Amount)
	}

	ExpectedProceedsComDate := StrToDate("2021-07-22 00:00:00")
	if OMSProceedPersisterObject[0].Date != ExpectedProceedsComDate {
		t.Errorf("TestProceeds Amount: Esperado (%s), Recebido (%s)", ExpectedProceedsComDate.String(), OMSProceedPersisterObject[0].Date.String())
	}

	ExpectedProceedsEvent := "JRS CAP PROPRIO"
	if OMSProceedPersisterObject[0].Event != ExpectedProceedsEvent {
		t.Errorf("TestProceeds Amount: Esperado (%s), Recebido (%s)", ExpectedProceedsEvent, OMSProceedPersisterObject[0].Event)
	}
}
