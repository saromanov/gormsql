package core

// Table represents gorm model definition
type Table struct {
	Name    string
	Columns map[string]Column
}

type Column struct {
	Name        string
	Annotations string
	Type        string
	PrimaryKey  bool
}
