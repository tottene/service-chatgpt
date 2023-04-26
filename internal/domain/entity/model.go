package entity

type Model struct {
	Name      string
	MaxTokens int
}

func NewModel(aName string, aMaxtokens int) *Model {
	return &Model{
		Name:      aName,
		MaxTokens: aMaxtokens,
	}
}

func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}
