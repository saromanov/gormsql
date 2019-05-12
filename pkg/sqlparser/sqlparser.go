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

// Parse provides parsing of sql query
func Parse(s string) (*core.Table, error) {
	table := &core.Table{}
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
		table.Name = name
		columns := []core.Column{}
		for _, c := range result.TableSpec.Columns {
			columns = append(columns, core.Column{
				Name:        c.Name.String(),
				Type:        c.Type.Type,
				Annotations: consuructColumnAnnotation(c.Type),
			})
		}
		table.Columns = columns
	}
	return table, nil
}

// constructColumnAnnotation adds gorm annotations to the model
func consuructColumnAnnotation(c sqlparser.ColumnType) string {
	response := "`gorm:`"
	if c.NotNull {
		response += "NOT NULL;"
	}
	if c.Autoincrement {
		response += "AUTOINCREMENT;"
	}
	if c.Default != nil {
		response += fmt.Sprintf(`DEFAULT:"%s"`, string(c.Default.Val))
	}
	return response + "`"
}
