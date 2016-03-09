package main

import (
	"fmt"
	"time"
)

func main() {
	speak("hello")
}

func speak(word string) {
	for i := 0; i < 5; i++ {
		fmt.Println(word)
		time.Sleep(300)
	}
}
