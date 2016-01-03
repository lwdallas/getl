package stream_test

import (
	"stream"
	"testing"
)

func TestBasicStream( t *testing.T ) {
	const control, data, error = "Control", "Data", "Error"
	var stream stream.Stream
	stream.SetControlMessage("Control")
	stream.SetDataMessage("Data")
	stream.SetErrorMessage("Error")

	if stream.ControlMessage != control || stream.DataMessage != data || stream.ErrorMessage != error {
		t.Errorf( "Stream object failed basic creation." )
	}

}
