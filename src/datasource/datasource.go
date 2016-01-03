package datasource

/*
	A data source contains information about the source of data for a row or a field
*/

type DataSource struct {
	Type   string
	Name   string
	Table  string
	Schema string
}

func init() {
}
