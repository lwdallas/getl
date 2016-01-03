package field_test

import (
	"field"
	"testing"
)

func TestBasicFieldObjectCreation( t *testing.T ) {
	const name, value, nullval = "Name", "Value", false
	var tfield field.Field
	tfield.Name = "Name"
	tfield.Value = "Value"
	tfield.Null = false

	if tfield.Name != name || tfield.Value != value || tfield.Null != nullval {
		t.Errorf( "Field failed basic field creation" )
	}

	var tfieldMetaData field.FieldMetaData
	tfieldMetaData.Type = "String"
	tfield.MetaData = tfieldMetaData

	const typeval = "String"

	if tfield.MetaData.Type != typeval {
		t.Errorf( "field metadata failed assignment" )
	}

}
