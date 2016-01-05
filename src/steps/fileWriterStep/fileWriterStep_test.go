package fileWriterStep_test

import (
	"steps/fileReaderStep"
	"steps/fileWriterStep"
	"testing"
)

func TestStepCreation(t *testing.T) {
	const idval, nameval = 1, "TestReaderStep"
	var trstep fileReaderStep.FileReaderStep

	trstep.Initialize(nil)
	trstep.ID = 1
	trstep.Name = "TestReaderStep"
	if trstep.ID != idval || trstep.Name != nameval {
		t.Errorf("basic reader step creation failed assignment")
	}

	trstep.Execute()
	trstep.Print()

	var twstep fileWriterStep.FileWriterStep

	twstep.Initialize(nil)
	twstep.ID = 2
	twstep.Name = "TestWriterStep"
	const idval2, nameval2 = 2, "TestWriterStep"
	if twstep.ID != idval2 || twstep.Name != nameval2 {
		t.Errorf("basic writer step creation failed assignment")
	}

	twstep.Execute()
	twstep.Print()

}
