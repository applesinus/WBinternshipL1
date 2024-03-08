package main

import "fmt"

// Creating a struct Human with some fields
type Human struct {
	Name string
	Age  int
}

// Creating method Walk()
func (person Human) Walk() {
	fmt.Printf("%s is walking\n", person.Name)
}

// Creating method Run()
func (person Human) Run() {
	fmt.Printf("%s is running\n", person.Name)
}

// Creating a struct Action
type Action struct {
	// Embedding struct Human into struct Action
	Human
}

func Ex1() {
	fmt.Printf("\n==========  Exercise 1:  ==========\n\n")

	fmt.Printf("Human:\n")
	// Creating a Human object
	person := Human{"Dmitry", 22}
	// Calling methods
	person.Walk()
	person.Run()

	fmt.Printf("\nAction:\n")
	// Creating an Action object and embedding Human into it
	action := Action{person}
	// Calling methods
	action.Walk()
	action.Run()

	fmt.Printf("\n===================================\n\n")
}
