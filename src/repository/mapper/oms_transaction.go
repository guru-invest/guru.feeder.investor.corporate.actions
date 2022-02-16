package mapper

//TODO - Remover comentários após validação

type OMSTransaction struct {
	//Dados de transações
	ID       int
	Symbol   string
	Quantity int
	Price    float64

	//Dados calculados com base nos eventos corporativos
	PostEventQuantity int
	PostEventPrice    float64

	// Dados do evento corporativo que estão na Financial
	PostEventSymbol string
	Factor          float64
	Event           string
}
