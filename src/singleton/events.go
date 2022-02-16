package singleton

type Events struct {
	Update    string
	Unfolding string
	Grouping  string
}

var instance *Events = nil

func New() *Events {
	if instance == nil {
		instance = new(Events)
		instance.Update = "ATUALIZACAO"
		instance.Unfolding = "DESDOBRAMENTO"
		instance.Grouping = "GRUPAMENTO"
	}
	return instance
}
