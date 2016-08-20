package main

import (
	"fmt"

	"./communication"
)
import "./snake"

const hostName string = "snake.cygni.se"
const port int = 80
const mode string = "training"

func main() {
	fmt.Println("main start")
	c := communication.NewClient(hostName, port, mode)
	defer c.Close()

	s := snake.NewSnake("golangsnake", c)
	fmt.Println("initiating snake")
	s.Init()

	var message string
	select {
	case message = <-s.FinishChannel:
		fmt.Println("Snake finished: ", message)
		return
	}
}
