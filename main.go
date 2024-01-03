package main

import (
	"log"
)

func main() {
	log.Println(Add(1, 2))
}


func Add(i int, j int) int {
	return i + j
}