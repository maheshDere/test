package check

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Url struct {
	Url    string
	Status bool
}

func Check2(myUrl Url, wg *sync.WaitGroup) {
	timeout := time.Duration(4 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(myUrl.Url)
	if err != nil {
		myUrl.Status = false
	}
	fmt.Printf("url: %s  status: %v \n", myUrl.Url, myUrl.Status)
	wg.Done()
}
