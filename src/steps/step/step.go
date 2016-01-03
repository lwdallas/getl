package step

import (
	"log"
	"stream"
)

type Step struct {
	Name    string
	ID      int
	Channel chan stream.Stream
	Data    string
}

func (self *Step) Initialize(messaging_chan chan stream.Stream) {
	log.Println("Initialized Step")
	log.Printf("Step config: %#v \n", self)
	self.Channel = messaging_chan
}

func (self *Step) Print() {
	log.Printf("%#v \n", self)
}

func (self *Step) SendDataMessageString(aValue string) {
	// send data down the stream
	var aStream stream.Stream
	aStream.DataMessage = aValue
	self.Channel <- aStream
}

func (self *Step) SendDataMessage() {
	// send data down the stream
	var aStream stream.Stream
	aStream.DataMessage = self.Data
	self.Channel <- aStream
}

func (self *Step) SendErrorMessage(aValue string) {
	// send an error down the stream
	var aStream stream.Stream
	aStream.ErrorMessage = aValue
	self.Channel <- aStream
}

func (self *Step) SendControlMessage(aValue string) {
	// send a control message down the stream
	var aStream stream.Stream
	aStream.ControlMessage = aValue
	self.Channel <- aStream
}

func (self *Step) ReadDataMessage() string {
	// read a data message from up the stream
	var aStream stream.Stream

	return (aStream.DataMessage)
}

func (self *Step) ReadErrorMessage() string {
	// read an error message from up the stream
	var aStream stream.Stream

	return (aStream.ErrorMessage)
}

func (self *Step) ReadControlMessage() string {
	// read an error message from up the stream
	var aStream stream.Stream

	return (aStream.ControlMessage)
}
