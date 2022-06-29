package main

import (
	"log"
	"time"

	"github.com/eisenwinter/go-panandtilt/driver"
)

func main() {
	pt, err := driver.Initialize(500 * time.Millisecond)
	if err != nil {
		log.Panic(err)
	}
	defer pt.Close()
	for {
		pt.Pan(-90)
		time.Sleep(3 * time.Second)
		pt.Pan(0)
		time.Sleep(3 * time.Second)
		pt.Pan(90)
		time.Sleep(3 * time.Second)
		pt.Pan(0)
		time.Sleep(3 * time.Second)
	}
}
