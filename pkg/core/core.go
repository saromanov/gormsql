// Package core defines main logic
package core

type Core struct {
	table Table
}

// New creates core object
func New(t Table) *Core {
	return &Core{
		table: t,
	}
}
