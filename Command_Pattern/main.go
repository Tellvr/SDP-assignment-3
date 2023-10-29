package main

import "fmt"

type Command interface {
	Execute()
}

type Build struct {
	Material string
}

func (b *Build) Execute() {
	fmt.Printf("Building with %s\n", b.Material)
}

type Demolish struct {
	Method string
}

func (d *Demolish) Execute() {
	fmt.Printf("Demolishing using %s\n", d.Method)
}

func main() {

	// Initialize commands
	build := &Build{Material: "concrete"}
	demolish := &Demolish{Method: "explosives"}

	// Execute commands
	build.Execute()
	demolish.Execute()

}
