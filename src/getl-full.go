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

func worker(messaging_chan, done_chan chan stream.Stream) {

	fmt.Println("...in worker...")

	var s sampleStep.SampleStep
	s.Initialize(messaging_chan)
	s.ID = 1
	s.Name = "Main Simple Step Test"
	s.Print()
	s.Data = "Data From Main Simple Step Test"
	s.SendDataMessage()
	s.Print()
}

func main() {
	//var messaging_chan chan stream.Stream
	//var shutdown_chan chan stream.Stream

	fmt.Println("Starting with all GETL modules")

	// test a stream
	fmt.Println("Testing a simple stream object")
	var aStream stream.Stream
	aStream.SetControlMessage("Control")
	aStream.SetDataMessage("Data")
	aStream.SetErrorMessage("Error")
	aStream.Print()

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

	messaging_chan := make(chan stream.Stream)
	shutdown_chan := make(chan stream.Stream)

	go worker(messaging_chan, shutdown_chan)

	select {
		case aStream = <-messaging_chan:
			fmt.Println("Inspecting current state of stream channel")
			fmt.Println(aStream)
	}

	fmt.Println("Finished")
}
