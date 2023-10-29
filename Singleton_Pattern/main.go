package main

import (
	"fmt"
	"time"
)

func main() {

	go handleCall("Call 1")
	go handleCall("Call 2")

	time.Sleep(2 * time.Second)

}

func handleCall(callNumber string) {
	assistant := GetInstance()

	fmt.Printf("Handling %s\n", callNumber)
	assistant.HandleCall()
}
