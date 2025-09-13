package main

import (
	"fmt"
)

type getRequest struct {
	reply chan int
}
type setRequest struct {
	value int
}

var (
	getChan = make(chan getRequest)
	setChan = make(chan setRequest)
)

func SetState(val int) {
	setChan <- setRequest{value: val}
}

func GetState() int {
	reply := make(chan int)
	getChan <- getRequest{reply: reply}
	return <-reply
}

func main() {

	go func() {
		state := 0
		for {
			select {
			case req := <-setChan:
				state = req.value
			case req := <-getChan:
				req.reply <- state
			}
		}
	}()

	SetState(5)
	fmt.Println("State:", GetState())

	SetState(10)
	fmt.Println("State:", GetState())
}
