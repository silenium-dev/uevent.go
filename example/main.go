package main

import (
	"log"

	"github.com/silenium-dev/uevent.go"
)

func main() {
	r, err := uevent.NewReader()
	if err != nil {
		log.Fatal(err)
	}

	dec := uevent.NewDecoder(r)

	msgChan := make(chan *uevent.Uevent, 10000)
	go func() {
		for evt := range msgChan {
			log.Printf("%+v", evt)
		}
	}()

	for {
		evt, err := dec.Decode()
		if err != nil {
			log.Fatal(err)
		}
		msgChan <- evt
	}
}
