package fileReaderStep

import (
	"bufio"
	"log"
	"os"
	"steps/step"
	"stream"
)

type FileReaderStep struct {
	step.Step
}

func (self *FileReaderStep) Initialize(messaging_chan chan stream.Stream) {
	log.Println("Initialized SampleStep")
	log.Printf("Step config: %#v \n", self)
	self.Channel = messaging_chan
}

func (self *FileReaderStep) Execute() {
	const fn = "/Users/lonniewebb/Documents/Projects/getl/src/steps/fileReaderStep/test-data/sample.txt"
	f, err := os.Open(fn)
	if err != nil {
		log.Println("error reading file ", err)
	} else {

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var aStream stream.Stream
			aStream.SetDataMessage(scanner.Text())
		}

	}
}
