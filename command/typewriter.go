package command

import "fmt"

// Printer defines the printing interface of the receiver.
type Printer interface {
	Print()
}

// CommandInvoker defines the invoking interface.
type CommandInvoker interface {
	Press(Command) CommandInvoker
	UndoLast() CommandInvoker
}

// Typewriter represents the command receiver with invoker methods.
type Typewriter struct {
	Draft          string
	Controller     *TypewriterController
	commandHistory []Command
}

// Press implements the execution invoker method with adding to command history.
func (t *Typewriter) Press(cmd Command) CommandInvoker {
	cmd.Execute()
	t.commandHistory = append(t.commandHistory, cmd)
	return t
}

// UndoLast implements the command invoker undo action from command history.
func (t *Typewriter) UndoLast() CommandInvoker {
	historyLen := len(t.commandHistory)
	if historyLen == 0 {
		return t
	}
	t.commandHistory[historyLen-1].Undo()
	t.commandHistory = t.commandHistory[:historyLen-1]
	return t
}

// Print implements Printer interface.
func (t *Typewriter) Print() {
	fmt.Print(t.Draft)
}

// NewTypewriter constructs a new Typewriter.
func NewTypewriter() *Typewriter {
	tw := &Typewriter{Draft: ""}
	tw.Controller = NewTypewriterController(tw)
	return tw
}

// Interface implementation assertion.
var (
	_ Printer        = &Typewriter{}
	_ CommandInvoker = &Typewriter{}
)
