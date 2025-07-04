package command

// Command defines the command interface for the actions.
type Command interface {
	Undo()
	Execute()
}

// PressCharacter implements a single character key press command.
type PressCharacter struct {
	value    rune
	receiver *Typewriter
}

// Key implements the single character value setter for the command.
func (p *PressCharacter) Key(char rune) *PressCharacter {
	p.value = char
	return p
}

// Execute implements the execution method of the command.
func (p *PressCharacter) Execute() {
	p.receiver.Controller.PressCharacter(p.value)
}

// Undo implements the undoing method of the command.
func (p *PressCharacter) Undo() {
	p.receiver.Controller.PressBackspace()
}

// NewPressCharacter constructs the instance of the PressCharacter command.
func NewPressCharacter(receiver *Typewriter) *PressCharacter {
	return &PressCharacter{receiver: receiver}
}

// PressBackspace implements backspace key press command.
type PressBackspace struct {
	receiver *Typewriter
	captured rune
}

// Execute implements the execution method of the command.
func (p *PressBackspace) Execute() {
	p.captured = p.receiver.Controller.PressBackspace()
}

// Undo implements the undoing method of the command.
func (p *PressBackspace) Undo() {
	if p.captured != 0 {
		p.receiver.Controller.PressCharacter(p.captured)
	}
}

// NewPressBackspace constructs the instance of the PressBackspace command.
func NewPressBackspace(receiver *Typewriter) *PressBackspace {
	return &PressBackspace{receiver: receiver}
}

// PressTab implements the Tab key press command.
type PressTab struct {
	receiver *Typewriter
	tabSize  byte
}

// Execute implements the execution method of the command.
func (p *PressTab) Execute() {
	p.receiver.Controller.PressTab(p.tabSize)
}

// Undo implements the undoing method of the command.
func (p *PressTab) Undo() {
	for range p.tabSize {
		p.receiver.Controller.PressBackspace()
	}
}

// NewPressTab constructs the instance of the PressTab command.
func NewPressTab(receiver *Typewriter, tabSize byte) *PressTab {
	return &PressTab{receiver: receiver, tabSize: tabSize}
}

// PressSpace implements the spacebar press command.
type PressSpace struct {
	receiver *Typewriter
}

// Execute implements the execution method of the command.
func (p *PressSpace) Execute() {
	p.receiver.Controller.PressSpace()
}

// Undo implements the undoing method of the command.
func (p *PressSpace) Undo() {
	p.receiver.Controller.PressBackspace()
}

// NewPressSpace constructs the instance of the PressSpace command.
func NewPressSpace(receiver *Typewriter) *PressSpace {
	return &PressSpace{receiver: receiver}
}

// PressEnter implements the Enter key press command.
type PressEnter struct {
	receiver *Typewriter
}

// Execute implements the execution method of the command.
func (p *PressEnter) Execute() {
	p.receiver.Controller.PressEnter()
}

// Undo implements the undoing method of the command.
func (p *PressEnter) Undo() {
	p.receiver.Controller.PressBackspace()
}

// NewPressEnter constructs the instance of the PressEnter command.
func NewPressEnter(receiver *Typewriter) *PressEnter {
	return &PressEnter{receiver: receiver}
}

// Interface implementation assertion.
var (
	_ Command = &PressCharacter{}
	_ Command = &PressSpace{}
	_ Command = &PressBackspace{}
	_ Command = &PressTab{}
	_ Command = &PressEnter{}
)
