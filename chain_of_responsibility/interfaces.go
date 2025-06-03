package chain_of_responsibility

import "fmt"

// Chainable defines the interface for a processing step in the chain of responsibility.
// Each step processes part of a meal and delegates to the next step, if any.
type Chainable interface {
	Nestable
	Handle() bool
	process()
}

// Nestable provides accessors for chaining steps together.
type Nestable interface {
	AddNext(step Chainable) Chainable
	getNext() Chainable
}

// Layerable defines how ingredients or modifications are added to a meal.
type Layerable interface {
	AddLayer(string)
}

// Combinable declares the interface of a completed meal that can accumulate layers and could be rendered as a string.
type Combinable interface {
	fmt.Stringer
	Layerable
	combineLayers() string
}
