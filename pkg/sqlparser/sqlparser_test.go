package sqlparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	_, err := Parse("aaa")
	assert.Error(t, err)

	_, err = Parse("INSERT")
	assert.Error(t, err)

	res, err := Parse(`
	CREATE TABLE IF NOT EXISTS app_user (
		id int NOT NULL,
		username varchar(45) NOT NULL,
	  );
	`)
	assert.NoError(t, err)
	_, ok := res["app_user"]
	if !ok {
		t.Errorf("should contains app_user")
	}
	_, ok = res["id"]
	if !ok {
		t.Errorf("should contains username")
	}
	assert.NoError(t, err)
}
