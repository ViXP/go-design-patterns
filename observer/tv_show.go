package observer

import (
	"fmt"
	"sync"
)

// TvShow represents a single broadcastable TV show.
type TvShow struct {
	title        string
	broadcasters []Observer
}

// Broadcast sends a message to all broadcasters subscribed to this TV show.
func (t *TvShow) Broadcast(message string) {
	t.Notify(fmt.Sprintf("[%v]: %v", t.title, message))
}

// Notify implements the Observable interface by publishing updates to all subscribed broadcasters.
func (t *TvShow) Notify(data any) {
	wg := sync.WaitGroup{}
	for _, b := range t.broadcasters {
		wg.Add(1)
		go func(b Observer) {
			defer wg.Done()
			b.Update(data)
		}(b)
	}
	wg.Wait()
}

// Subscribe adds a broadcaster to the list of subscribers for this TV show.
func (t *TvShow) Subscribe(listener Observer) {
	t.broadcasters = append(t.broadcasters, listener)
}

// Unsubscribe allows to unsubscribe broadcaster from the TV show.
func (t *TvShow) Unsubscribe(listener Observer) {
	var index int = -1

	for i, b := range t.broadcasters {
		if b == listener {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("this broadcaster wasn't subscribe to this show before")
		return
	}

	t.broadcasters = append(t.broadcasters[:index], t.broadcasters[index+1:]...)
}

// NewTvShow creates the new TV show.
func NewTvShow(title string) *TvShow {
	return &TvShow{title: title}
}

// Interface implementation assertion.
var _ Broadcasting = &TvShow{}
