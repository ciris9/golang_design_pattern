package main

import (
	"fmt"
	"sync"
)

// todo golang sync.once
// sync.once 可以保证Do内的函数，永远只被执行一次
var once = &sync.Once{}
var singleInstanceWithOnce *single

func getInstanceWithOnce() *single {
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		singleInstanceWithOnce = &single{}
	})
	return singleInstanceWithOnce
}

// todo don't use golang sync.once
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		}
	}

	return singleInstance
}

func main() {
	fmt.Println("sync.once")
	getInstanceWithOnce()
	getInstanceWithOnce()
	getInstanceWithOnce()
	fmt.Println("no sync.once")
	getInstance()
	getInstance()
	getInstance()
	/*
		result:
		sync.once
		Creating single instance now.
		no sync.once
		Creating single instance now.
	*/
}
