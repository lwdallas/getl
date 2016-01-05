package fileWriterStep

import (
	"bufio"
	"log"
	"os"
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

func (self *FileWriterStep) Execute() {
	const fn = "/Users/lonniewebb/Documents/Projects/getl/src/steps/fileReaderStep/test-data/sample.out"
	f, err := os.Open(fn)
	if err != nil {
		f, err = os.Create(fn)
	}

	if err != nil {
		log.Println("error opening file ", err)
	} else {
		defer f.Close()

		// TODO need to pop a message here

		w := bufio.NewWriter(f)
		_, err := w.WriteString("test line")
		if err != nil {
			log.Println("Cannot write files")
		}
		w.Flush()

		f.Sync()
		f, _ = os.Open(fn)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			log.Println(scanner.Text())
		}

	}
}
