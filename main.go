package main

import (
	ws "TestExample/website"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Application Starting ...\n")

	websites := []ws.Website{
		{"https://www.google.com", "Google", false},
		{"https://www.facebook.com", "Facebook", false},
		{"https://twitter.com", "Twitter", false},
		{"https://pradipkshirsagar.com", "PradipInfo", false},
		{"https://youtube.com", "YouTube", false},
	}

	wg := sync.WaitGroup{}

	var ch = make(chan ws.Website)

	for _, p := range websites {
		wg.Add(1)
		w := p
		go w.Check(&wg, ch)
		w = <-ch
		fmt.Print(w.Url, " is ")
		if w.Status {
			fmt.Println("Up")
		} else {
			fmt.Println("Down")
		}
	}
	wg.Wait()
	fmt.Println("\nApplication Ends ...")
}
