package oms

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyBasicCorporateAction(OMSTransaction mapper.OMSTransaction, corporate_action mapper.CorporateAction) mapper.OMSTransaction {
	targetTicket := corporate_action.TargetTicker
	if corporate_action.TargetTicker == "" {
		targetTicket = corporate_action.Symbol
	}
	OMSTransaction.EventFactor = corporate_action.CalculatedFactor
	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if corporate_action.Description == constants.Update {
		OMSTransaction.EventFactor = 1
	}

	OMSTransaction.EventName = corporate_action.Description
	OMSTransaction.PostEventSymbol = targetTicket
	OMSTransaction.EventDate = corporate_action.ComDate
	OMSTransaction.PostEventQuantity = OMSTransaction.PostEventQuantity / OMSTransaction.EventFactor
	OMSTransaction.PostEventPrice = utils.Truncate(OMSTransaction.Amount/OMSTransaction.PostEventQuantity, 2)

	// Processo cumulativo
	// OMSTransaction.Quantity = int(utils.Truncate(OMSTransaction.PostEventQuantity, 0))
	// OMSTransaction.Price = OMSTransaction.PostEventPrice

	return OMSTransaction
}
