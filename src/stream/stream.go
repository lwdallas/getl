package stream

import "fmt"

/*
	controlMessage string - control messages that determine process flow
	dataMessage string - data and meta data concerning fields in the stream and references to data sources. Not data sources themselves
	errorMessage string - error state messages
*/

type Stream struct {
	ControlMessage string
	DataMessage    string
	ErrorMessage   string
}

func (self *Stream) Print() {
	fmt.Printf("Self control message: %#v", self.ControlMessage)
	fmt.Printf("Self data message: %#v", self.DataMessage)
	fmt.Printf("Self error message: %#v", self.ErrorMessage)
}

func init() {}

func (self *Stream) SetControlMessage(aValue string) error {
	self.ControlMessage = aValue
	return nil
}

func (self *Stream) SetDataMessage(aValue string) error {
	self.DataMessage = aValue
	return nil
}

func (self *Stream) SetErrorMessage(aValue string) error {
	self.ErrorMessage = aValue
	return nil
}
