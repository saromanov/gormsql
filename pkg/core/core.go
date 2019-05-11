// Package core defines main logic
package core

import (
	"os"
	"reflect"

	"github.com/pkg/errors"
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
	name := c.table.Name
	if name == "" {
		return errNoName
	}

	f, err := os.Create(name + ".go")
	if err != nil {
		return errors.Wrap(err, "unable to create file")
	}

	res, err := c.generate()
	if err != nil {
		return errors.Wrap(err, "unable to generate model")
	}

	_, err = f.Write([]byte(res))
	if err != nil {
		return errors.Wrap(err, "unable to write to file")
	}
	defer f.Close()
	return nil
}

// generate provides generating of the model
func (c *Core) generate() (string, error) {
	ref := []reflect.StructField{}
	for _, col := range c.table.Columns {
		v := reflect.StructField{
			Name: col.Name,
		}
		if col.Annotations != "" {
			v.Tag = reflect.StructTag(col.Annotations)
		}
		switch col.Type {
		case "varchar":
			v.Type = reflect.TypeOf(string(""))
		case "initeger":
			v.Type = reflect.TypeOf(int(0))
		}
		ref = append(ref, v)
	}
	return reflect.StructOf(ref).String(), nil
}
