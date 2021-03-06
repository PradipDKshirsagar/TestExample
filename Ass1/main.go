package main

import (
	ws "TestExample/Ass1/website"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Application Starting ...\n")

	websites := []ws.Website{
		{"https://www.google.com", "Google", false},
		{"https://www.facebook.com", "Facebook", false},
		{"https://twitter.com", "Twitter", false},
		{"https://youtube.com", "YouTube", false},
		{"https://pradipkshirsagar.com", "PradipInfo", false},
	}

	wg := sync.WaitGroup{}

	var ch = make(chan ws.Website)

	for _, w := range websites {
		wg.Add(2)
		go func(w ws.Website, ch chan<- ws.Website) {
			w.Check(&wg,ch)
		}(w, ch)

		go func(ch <-chan ws.Website) {
			w = <-ch
			fmt.Print(w.Url, " is ")
			if w.Status {
				fmt.Println("Up")
			} else {
				fmt.Println("Down")
			}
			fmt.Println("Done")
			wg.Done()
		}(ch)
	}
	wg.Wait()
	fmt.Println("\nApplication Ends ...")
}
