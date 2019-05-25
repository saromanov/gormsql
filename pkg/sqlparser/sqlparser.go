// Package sqlparser provides parsing of the input sql queries
package sqlparser

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/saromanov/gormsql/pkg/core"
	"github.com/xwb1989/sqlparser"
)

var errNotSupported = errors.New("not supported operation")

// Index defines helpful struct for index definition
type Index struct {
	PrimaryKey bool
}

// Parse provides parsing of sql query
func Parse(s string) (map[string]*core.Table, error) {
	indexes := map[string]Index{}
	tables := map[string]*core.Table{}
	stmt := sqlparser.NewStringTokenizer(s)
	for {
		stmt, err := sqlparser.ParseNext(stmt)
		if err == io.EOF {
			break
		}
		if stmt == nil {
			break
		}

		result := stmt.(*sqlparser.DDL)
		if result.Action != sqlparser.CreateStr {
			return nil, errNotSupported
		}
		name := result.Table.Name.String()
		if name == "" {
			name = result.NewName.Name.String()
		}
		if result.TableSpec == nil {
			return nil, errors.Wrap(err, fmt.Sprintf("unexpected error on the table definition: %s", name))
		}

		for _, i := range result.TableSpec.Indexes {
			if i.Info.Primary {
				for _, idx := range i.Columns {
					_, ok := indexes[idx.Column.String()]
					if !ok {
						indexes[idx.Column.String()] = Index{
							PrimaryKey: true,
						}
					}
				}
			}
		}
		table := &core.Table{Name: name}
		columns := map[string]core.Column{}
		for _, c := range result.TableSpec.Columns {
			idx, _ := indexes[c.Name.String()]
			columns[c.Name.String()] = core.Column{
				Name:        c.Name.String(),
				Type:        c.Type.Type,
				Annotations: consuructColumnAnnotation(c.Type, idx),
			}
		}
		table.Columns = columns
		tables[name] = table
	}
	if len(tables) == 0 {
		return nil, fmt.Errorf("unable to generate data")
	}
	return tables, nil
}

// constructColumnAnnotation adds gorm annotations to the model
func consuructColumnAnnotation(c sqlparser.ColumnType, idx Index) string {
	startGorm := "`gorm:`"
	response := startGorm
	if c.NotNull {
		response += "NOT NULL;"
	}
	if c.Autoincrement {
		response += "AUTOINCREMENT;"
	}
	if idx.PrimaryKey {
		response += "PRIMARY_KEY;"
	}
	if c.Default != nil {
		response += fmt.Sprintf("DEFAULT:`%s`", string(c.Default.Val))
	}
	if c.Length != nil && c.Default != nil {
		response += fmt.Sprintf(`SIZE:"%s"`, string(c.Default.Val))
	}
	if response == startGorm {
		return ""
	}
	return response + "`"
}
