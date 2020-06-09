package main

import "fmt"

func main() {
	greeting := "Hello Go!"
	fmt.Println(greeting)

	prt := &greeting
	fmt.Println(prt, *prt)

	greeting = "Hello world from Go!"
	fmt.Println(prt, *prt)
}
