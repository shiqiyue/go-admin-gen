package dto

type Model struct {
	Name string

	Description string

	Fields []*ModelField
}

type ModelField struct {
	Name string

	Description string

	Type string

	Ptr bool

	Tag string
}
