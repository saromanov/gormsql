// Package core defines main logic
package core

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
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
			Name: strings.Title(col.Name),
		}
		if col.Annotations != "" {
			v.Tag = reflect.StructTag(col.Annotations)
		}
		switch col.Type {
		case "varchar":
			v.Type = reflect.TypeOf(string(""))
		case "integer", "int":
			v.Type = reflect.TypeOf(int(0))
		case "timestamp":
			v.Type = reflect.TypeOf(time.Time{})
		}
		ref = append(ref, v)
	}
	return fmt.Sprintf("type %s %s", strings.Title(c.table.Name), reflect.StructOf(ref).String()), nil
}
