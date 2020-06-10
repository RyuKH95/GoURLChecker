package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	go sexyCount("nico")
	sexyCount("flynn")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy ", i)
		time.Sleep(time.Second)
	}
}
