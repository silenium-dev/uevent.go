package main

import (
	"fmt"
	"log"

	"github.com/silenium-dev/uevent.go"
)

func main() {
	r, err := uevent.NewReader()
	if err != nil {
		log.Fatal(err)
	}

	dec := uevent.NewDecoder(r)

	for {
		evt, err := dec.Decode()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(evt)
	}
}
