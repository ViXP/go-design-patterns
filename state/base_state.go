package state

// BaseState provides the default behavior shared by all concrete States. It acts as an abstract foundation — concrete
// States can override only the actions they support while inheriting the default ones.
type BaseState struct {
	subject Context
}

// String should be overridden by each concrete State to return its name.
func (b *BaseState) String() string {
	panic("unimplemented")
}

// Return defines the default behavior for the Return action in all States.
func (b *BaseState) Return() {
	panic("This ticket cannot be returned in the current state!")
}

// Buy defines the default behavior for the Buy action in all States.
func (b *BaseState) Buy() {
	panic("This ticket has already been purchased.")
}

// CheckIn defines the default behavior for the CheckIn action in all States.
func (b *BaseState) CheckIn() {
	panic("This ticket has already been used for entry.")
}

// Reserve defines the default behavior for the Reserve action in all States.
func (b *BaseState) Reserve() {
	panic("This ticket cannot be reserved in the current state.")
}

// Use defines the default behavior for the Use action in all States.
func (b *BaseState) Use() {
	panic("This ticket cannot be used in the current state.")
}

// Interface implementation assertion.
var _ State = &BaseState{}
