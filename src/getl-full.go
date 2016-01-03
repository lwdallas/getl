package main

import (
	"datasource"
	"field"
	"fmt"
	"row"
	"stream"
)

func main() {
	fmt.Println("Starting with all GETL modules")

	// test a stream
	fmt.Println("Testing a simple stream object")
	var stream stream.Stream
	stream.SetControlMessage("Control")
	stream.SetDataMessage("Data")
	stream.SetErrorMessage("Error")
	stream.Print()

	// test a dataSource
	fmt.Println("Testing a simple datasource object")
	var dataSource datasource.DataSource
	fmt.Println(dataSource)

	// test a row
	fmt.Println("Testing a simple row object")
	var row row.Row
	fmt.Println(row)

	// test a field
	fmt.Println("Testing a simple field object")
	var field field.Field
	fmt.Println(field)

	fmt.Println("Finished")
}
