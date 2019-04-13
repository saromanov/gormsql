// Package sqlparser provides parsing of the input sql queries
package sqlparser

import (
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

// Parse provides parsing of sql query
func Parse(s string) error {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return errors.Wrap(err, "unable to parse sql query")
	}

	return nil
}
