package fileReaderStep_test

import (
	"steps/fileReaderStep"
	"testing"
)

func TestStepCreation(t *testing.T) {
	const idval, nameval = 1, "TestStep"
	var tstep fileReaderStep.FileReaderStep

	tstep.Initialize(nil)
	tstep.ID = 1
	tstep.Name = "TestStep"
	if tstep.ID != idval || tstep.Name != nameval {
		t.Errorf("basic step creation failed assignment")
	}

	tstep.Execute()
	tstep.Print()
}
