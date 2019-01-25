package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var done chan bool

func main() {
	done = make(chan bool)
	go count()

	cb := js.NewCallback(printStuff)
	defer cb.Release()

	d := js.Global().Get("navigator").Get("mediaDevices").Get("getUserMedia")
	d.Invoke(cb)

	<-done
}

func count() {
	for i := 0; i < 100; i++ {

		time.Sleep(time.Second * 1)

		fmt.Println(i)
	}
	done <- true
}

func printStuff(args []js.Value) {
	fmt.Println(args)
	done <- true
}
