package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/eisenwinter/go-panandtilt/driver"
)

func main() {
	pt, err := driver.Initialize(2 * time.Second)
	if err != nil {
		log.Panic(err)
	}
	defer pt.Close()
	for {
		t := time.Now()
		a := int(math.Sin(float64(t.Unix()*2)) * 90)
		pt.Pan(a)
		pt.Tilt(a)
		fmt.Printf("%d", a)
		time.Sleep(5 * time.Millisecond)

	}
}
