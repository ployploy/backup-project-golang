package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	ploychan := make(chan bool)
	go ploy(ploychan)
	go play()
	ployResult := <-ploychan
	fmt.Println(ployResult)
	fmt.Println("end")
}

func ploy(c chan bool) {
	for index := 0; index < 5; index++ {
		time.Sleep(1 * time.Second)
		fmt.Println("ploy")
	}
	c <- true
}

func play() {
	for index := 0; index < 10; index++ {
		time.Sleep(1 * time.Second)
		fmt.Println("play")

	}

}
