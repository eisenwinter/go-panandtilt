package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eisenwinter/go-panandtilt/driver"
)

func main() {
	pt, err := driver.Initialize(2 * time.Second)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	angle1, err := pt.PanValue()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	angle2, err := pt.TiltValue()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("pan = %v - tilt = %v\r\n", angle1, angle2)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf(text)
		text = strings.TrimRight(text, "\n")
		if strings.HasPrefix(text, "p:") {
			num := strings.TrimPrefix(text, "p:")
			n, e := strconv.Atoi(num)
			if e == nil {
				pt.Pan(n)
			}
			pv, _ := pt.PanValue()
			fmt.Printf("reads: %v\r\n", pv)
		}

		if strings.HasPrefix(text, "t:") {
			num := strings.TrimPrefix(text, "t:")
			n, e := strconv.Atoi(num)
			if e == nil {
				pt.Tilt(n)
			}
			tv, _ := pt.TiltValue()
			fmt.Printf("reads: %v\r\n", tv)
		}
	}

}
