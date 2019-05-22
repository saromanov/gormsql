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

// Core defines main logic of the app
type Core struct {
	tables   map[string]*Table
	fileName string
	dirName  string
}

// New creates core object
func New(dirName, filename string, t map[string]*Table) *Core {
	return &Core{
		tables:   t,
		fileName: filename,
		dirName:  dirName,
	}
}

// Do provides generation of the gorm tables
func (c *Core) Do() error {
	if c.fileName == "" {
		return errNoName
	}

	f, err := os.Create(c.fileName + ".go")
	if err != nil {
		return errors.Wrap(err, "unable to create file")
	}

	result := fmt.Sprintf("package %s\n", c.fileName)
	for _, t := range c.tables {
		res, err := c.generate(*t)
		if err != nil {
			return errors.Wrap(err, "unable to generate model")
		}
		result += fmt.Sprintf("%s\n", res)
	}

	_, err = f.Write([]byte(result))
	if err != nil {
		return errors.Wrap(err, "unable to write to file")
	}
	defer f.Close()
	return nil
}

// generate provides generating of the model
func (c *Core) generate(tab Table) (string, error) {
	ref := []reflect.StructField{}
	for _, col := range tab.Columns {
		v := reflect.StructField{
			Name: strings.Title(col.Name),
		}
		if col.Annotations != "" {
			v.Tag = reflect.StructTag(col.Annotations)
		}
		switch col.Type {
		case "float":
			v.Type = reflect.TypeOf(float64(0))
		case "varchar":
			v.Type = reflect.TypeOf(string(""))
		case "integer", "int", "tinyint":
			v.Type = reflect.TypeOf(int(0))
		case "bigint":
			v.Type = reflect.TypeOf(int64(0))
		case "timestamp":
			v.Type = reflect.TypeOf(time.Time{})
		}
		ref = append(ref, v)
	}
	return fmt.Sprintf("type %s %s", strings.Title(tab.Name), reflect.StructOf(ref).String()), nil
}
