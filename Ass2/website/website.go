package website

import (
	"sync"
	"net/http"
//	"fmt"
)

//Website holds the basic website info like url, name
type Website struct {
	Url    string
	Name   string
	Status bool
}

func (w Website) Check(wg *sync.WaitGroup,ch chan<- Website) {
//	fmt.Println("degub", w.Name)
	_, err := http.Get(w.Url)
	
	if err != nil {
		w.Status = false
		ch <- w
	} else {
		w.Status = true
		ch <- w
	}
	wg.Done()
}
