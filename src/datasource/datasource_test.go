package datasource_test

import (
	"datasource"
	"testing"
)

func TestBasicFieldObjectCreation(t *testing.T) {
	const name, schema, table, typeval = "Name", "Schema", "Table", "Type"
	var tdataSource datasource.DataSource
	tdataSource.Name = "Name"
	tdataSource.Schema = "Schema"
	tdataSource.Table = "Table"
	tdataSource.Type = "Type"

	if tdataSource.Name != name || tdataSource.Schema != schema || tdataSource.Table != table || tdataSource.Type != typeval {
		t.Errorf("Field failed basic field creation")
	}
}
