package main

import "fmt"

type worker struct {
	name string
	age  int
	job  job
}

type job struct {
	name   string
	salary float32
}

func main() {
	person1 := worker{name: "Daniel",
		age: 20,
		job: job{name: "Programmer", salary: 228322}}

	var person2 worker = worker{name: "Cool guy",
		age: 14,
		job: job{name: "Designer", salary: 228321}}

	var person3 = worker{}
	person3.name = "Maks"
	person3.age = 18
	person3.job = job{name: "Student", salary: 0}

	fmt.Println(person1) // {Daniel 20 {Programmer 228322}}
	fmt.Println(person2) // {Cool guy 14 {Designer 228321}}
	fmt.Println(person3) // {Maks 18 {Student 0}}
}
