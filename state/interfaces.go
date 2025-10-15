package state

import "fmt"

// Actioner defines the common set of actions shared by Contexts and States.
type Actioner interface {
	Reserve()
	Buy()
	CheckIn()
	Use()
	Return()
}

// Stateful defines the methods required to change or retrieve the current state of a structure.
type Stateful interface {
	ChangeState(state Actioner)
	GetCurrentState() Actioner
}

// State is the common interface for all concrete States that can handle actions and provide a stringified
// representation of themselves.
type State interface {
	Actioner
	fmt.Stringer
}

// Context is the common interface for all concrete Contexts that use States to perform actions.
type Context interface {
	Stateful
	Actioner
}
