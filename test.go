package main

import "fmt"

type Person struct{
	name string
	id int
	likes []string
}

func main(){
	person := Person{
		"sly",2008,"golang","java"}
	fmt.Println(person)
}