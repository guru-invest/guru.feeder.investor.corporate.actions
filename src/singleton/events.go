package singleton

type Events struct {
	Update           string
	Unfolding        string
	Grouping         string
	InterestOnEquity string
	Dividend         string
	Income           string
}

var instance *Events = nil

func New() *Events {
	if instance == nil {
		instance = new(Events)
		instance.Update = "ATUALIZACAO"
		instance.Unfolding = "DESDOBRAMENTO"
		instance.Grouping = "GRUPAMENTO"
		instance.InterestOnEquity = "JRS CAP PROPRIO"
		instance.Dividend = "DIVIDENDO"
		instance.Income = "RENDIMENTO"
	}
	return instance
}
