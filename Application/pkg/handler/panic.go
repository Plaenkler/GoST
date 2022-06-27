package handler

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"
)

var Debug bool = false
var index int = 0
var mutex *sync.Mutex = &sync.Mutex{}

func HandlePanic(section string) {
	if err := recover(); err != nil {
		mutex.Lock()
		defer mutex.Unlock()

		log.Printf("[%v] [%v] interruption: %v", section, index, err)

		if Debug {
			log.Printf("[%v] [%v] stack trace:", section, index)
			fmt.Println(string(debug.Stack()))
		}

		index++
	}
}
