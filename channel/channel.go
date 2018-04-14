package main

import (
	"github.com/satori/go.uuid"
	"time"
	"fmt"
)

type Producer struct {

	OutChan chan string

}

func (p *Producer) produce() {

	for {
		time.Sleep(time.Second)
		p.OutChan <- uuid.Must(uuid.NewV4()).String()
	}

}

func main() {

	readChan := make(chan string, 20)
	prod := Producer{
		OutChan: readChan,
	}

	go prod.produce()

	for {
		c := <-readChan
		t := time.Now().Local().String()
		fmt.Println("New UUID:", t, c)
	}

}
