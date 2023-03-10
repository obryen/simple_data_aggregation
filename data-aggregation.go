// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	user := fetchUser()
	reschan := make(chan interface{}, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("Hello", user)

	go fetchMatch(user, reschan, wg)
	go fetchLikes(user, reschan, wg)
	wg.Wait()
	close(reschan)
	for resp := range reschan {
		fmt.Println("resp", resp)
	}

	end := time.Since(start)
	fmt.Println("time taken", end)
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Sammy"
}

func fetchLikes(userName string, reschan chan interface{}, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	reschan <- 15
	wg.Done()
}
func fetchMatch(userName string, reschan chan interface{}, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	reschan <- "Lydia"
	wg.Done()
}
