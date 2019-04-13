// Package sqlparser provides parsing of the input sql queries
package sqlparser

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

var errNotSupported = errors.New("not supported operation")

// Parse provides parsing of sql query
func Parse(s string) error {
	stmt := sqlparser.NewStringTokenizer(s)

	for {
		stmt, err := sqlparser.ParseNext(stmt)
		if err == io.EOF {
			break
		}

		fmt.Println(stmt)
	}
	return nil
}
