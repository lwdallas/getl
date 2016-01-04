package fileWriterStep

import (
	"log"
	"steps/step"
	"stream"
)

type FileWriterStep struct {
	step.Step
}

func (self *FileWriterStep) Initialize(messaging_chan chan stream.Stream) {
	log.Println("Initialized SampleStep")
	log.Printf("Step config: %#v \n", self)
	self.Channel = messaging_chan
}
