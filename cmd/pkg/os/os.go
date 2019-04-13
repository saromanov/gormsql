// Package os implements io operations with files
package os

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// ReadFile returns file data based of path
func ReadFile(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read file")
	}
	return dat, nil
}
