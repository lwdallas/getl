package step

import "log"

type Step struct {
	Name string
	ID   int
}

func (self *Step) Initialize(messaging *chan int) {
	log.Println("Initialized Step")
	log.Printf("Step config: %#v \n", self)
}

func (self *Step) Print() {
	log.Printf("%#v \n", self)
}
