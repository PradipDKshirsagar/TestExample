package website

import (
	"net/http"
	//"sync"
)

//Website holds the basic website info like url, name
type Website struct {
	Url    string
	Name   string
	Status bool
}

func (w Website) Check(ch chan Website) {
	_, err := http.Get(w.Url)
	if err != nil {
		w.Status = false
		ch <- w
	} else {
		w.Status = true
		ch <- w
	}
	//wg.Done()
}
