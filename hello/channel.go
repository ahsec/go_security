package main

import (
	"log"
	"time"
)

func process(doneChannel chan bool) {
	time.Sleep(time.Second * 6)
	doneChannel <- true
}

func main() {
	var doneChannel chan bool
	doneChannel = make(chan bool)

	//	tempBool := <-doneChannel
	//	log.Println(tempBool)

	go process(doneChannel)

	readyToExit := false

	for !readyToExit {
		select {
		case done := <-doneChannel:
			log.Println("Done message received", done)
			readyToExit = true
		default:
			log.Println("No done signal yet. waiting")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
