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

func newTransactionObject(symbol string, quantity int, price float64, trade_date time.Time) mapper.OMSTransaction {
	result := mapper.OMSTransaction{}
	result.Symbol = symbol
	result.Quantity = quantity
	result.Price = price
	result.TradeDate = trade_date
	return result
}

func newCorporateActionObject(description string, value float64, com_date time.Time) mapper.CorporateAction {
	result := mapper.CorporateAction{}
	result.Description = description
	result.Value = value
	result.ComDate = com_date
	return result
}

func TestProceeds(t *testing.T) {

	// Dados chave para identificação dos ativos por usuários
	Customer := "fzVzgo8b"

	// Mock de transações
	OMSTransactionObjects := []mapper.OMSTransaction{}
	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BIDI4", 3, 27.41, StrToDate("2021-07-20 11:40:26")))
	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BIDI4", 5, 26.68, StrToDate("2021-07-26 10:29:32")))
	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BIDI4", 29, 26.68, StrToDate("2021-07-26 10:29:32")))
	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BIDI4", 1, 25.69, StrToDate("2021-07-28 8:11:05")))
	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BIDI4", 1, 25.67, StrToDate("2021-07-28 8:13:40")))
	OMSTransactionObjects = append(OMSTransactionObjects, newTransactionObject("BIDI4", 20, 9.47, StrToDate("2022-01-03 11:34:38")))
	// Agrupamento das transações em um map com chave por usuário
	OMSTransactions := map[string][]mapper.OMSTransaction{}
	OMSTransactions[Customer] = OMSTransactionObjects

	// Mock de eventos corporativos
	CorporateActionObjects := []mapper.CorporateAction{}
	CorporateActionObjects = append(CorporateActionObjects, newCorporateActionObject("JRS CAP PROPRIO", 0.012084364, StrToDate("2021-07-22 0:00:00")))
	// Agrupamento dos eventos corporativos em um map com chave por symbol
	CorporateActions := map[string][]mapper.CorporateAction{}
	CorporateActions["BIDI4"] = CorporateActionObjects

	// Processamento
	OMSProceedPersisterObject := []mapper.OMSProceeds{}
	OMSProceedPersisterObject = append(OMSProceedPersisterObject, oms.ApplyProceedsCorporateAction(Customer, "BIDI4", OMSTransactions, CorporateActions)...)

	// Teste
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
