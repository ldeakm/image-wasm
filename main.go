package main

import (
	"fmt"
	"syscall/js"
	"time"
	"upcutil"

var done chan bool

func main() {
	done = make(chan bool)
	go count()

// What I need to do
// 1.) Get access to the video stream
// 2.) Process that in webassembly looking for a UPC
// 3.) If I one
// 3a.) Create an image of the UPC and display it in the camera frame
// 3b.) Get the deatils for the UPC
// 4.) Paint the resutls


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
