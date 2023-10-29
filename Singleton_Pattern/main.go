package main

import (
	"fmt"
	"sync"
	"time"
)

type CallCenterAssistant struct {
	Name string
}

var instance *CallCenterAssistant
var once sync.Once

func GetInstance() *CallCenterAssistant {
	once.Do(func() {
		instance = &CallCenterAssistant{
			Name: "Assistant 1",
		}
		fmt.Println("Initializing assistant...")
	})

	return instance
}

func (a *CallCenterAssistant) HandleCall() {
	fmt.Printf("Assistant %s handling call\n", a.Name)

	a.GreetCustomer()
	a.RecommendProduct()
	a.ThankCustomer()
}

func (a *CallCenterAssistant) GreetCustomer() {
	fmt.Println("Greeting customer...")
}

func (a *CallCenterAssistant) RecommendProduct() {
	fmt.Println("Recommending product...")
}

func (a *CallCenterAssistant) ThankCustomer() {
	fmt.Println("Thanking customer...")
}

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
