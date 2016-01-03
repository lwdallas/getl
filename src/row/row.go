package row

import (
	"datasource"
	"field"
)

/*
	A row is a collection of fields, an atomic element of a stream
*/

type Row struct {
	Fields      []*field.Field
	rowMetaData RowMetaData
}

type RowMetaData struct {
	dataSource *datasource.DataSource
}

func init() {
}
