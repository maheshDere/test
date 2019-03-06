package main

import (
	"fmt"
	"sync"
	"task/check"
)

func main() {
	var wg sync.WaitGroup
	var Site = [...]string{
		"http://google.com",
		"http://mwertsr.com",
		"http://yturet.com",
	}
	var len = len(Site)
	for i := 0; i < len; i++ {
		wg.Add(1)
		go check.Check2(check.Url{Site[i], true}, &wg)
	}
	wg.Wait()
	fmt.Println("main ended")
}
