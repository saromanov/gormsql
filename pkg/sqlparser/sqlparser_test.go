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
	table, ok := res["app_user"]
	if !ok {
		t.Errorf("should contains app_user")
	}
	if table == nil {
		t.Errorf("table is not exist")
	}
	assert.Equal(t, "app_user", table.Name)
	assert.Equal(t, 2, len(table.Columns))
}
