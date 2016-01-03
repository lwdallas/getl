package main

/*
	This is the full version of the app with all steps linked in
*/

import (
	"datasource"
	"field"
	"fmt"
	"row"
	"steps/sampleStep"
	_ "steps/step"
	"stream"
)

var messaging_chan chan int
var shutdown_chan chan int

func worker(messaging_chan, done_chan chan int) {

	fmt.Println("...in worker...")

	var s sampleStep.SampleStep
	s.Initialize(&messaging_chan)
	s.ID = 1
	s.Name = "Main Simple Step Test"
	s.Print()
	messaging_chan <- 0
}

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

	messaging_chan = make(chan int)
	shutdown_chan = make(chan int)

	go worker(messaging_chan, shutdown_chan)

	<-messaging_chan

	fmt.Println("Finished")
}
