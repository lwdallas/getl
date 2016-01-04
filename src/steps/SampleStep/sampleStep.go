package sampleStep

import (
	"log"
	"steps/step"
	"stream"
)

type SampleStep struct {
	step.Step
}

func (self *SampleStep) Initialize(messaging_chan chan stream.Stream) {
	log.Println("Initialized SampleStep")
	log.Printf("Step config: %#v \n", self)
	self.Channel = messaging_chan
}
