// Package core defines main logic
package core

import (
	"errors"
)

var errNoName = errors.New("name is not defined")

type Core struct {
	table Table
}

// New creates core object
func New(t Table) *Core {
	return &Core{
		table: t,
	}
}

// Do provides generation of the gorm tables
func (c *Core) Do() error {
	if c.table.Name == "" {
		return errNoName
	}
	return nil
}
