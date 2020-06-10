package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "guy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	// time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
