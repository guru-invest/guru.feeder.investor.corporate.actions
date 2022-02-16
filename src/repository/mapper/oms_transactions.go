package mapper

import "time"

//TODO - Remover comentários após validação

type OMSTransaction struct {
	//Dados de transações
	//TODO - De onde vem esses dados ?
	Symbol    string
	Quantity  int
	Price     float64
	Amount    float64
	TradeDate time.Time

	//Dados calculados com base nos eventos corporativos
	PostEventQuantity int
	PostEventPrice    float64

	// Dados do evento corporativo que estão na Financial
	PostEventSymbol string
	Factor          float64
	ComDate         time.Time
	Event           string
}
