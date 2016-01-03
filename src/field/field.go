package field

/*
	A field is a data structure that is streamed an atomic element of a row
*/

type Field struct {
	MetaData FieldMetaData
	Name     string
	Value    string
	Null     bool
}

type FieldMetaData struct {
	Type string
}

func init() {
}
