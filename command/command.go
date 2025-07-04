// Package command implements the Command design pattern.
// Typewriter acts as both the Invoker and Receiver in the Command pattern.
// It uses its own implementation of the CommandInvoker interface to execute
// each command sequentially, with the ability to undo every action.
// All commands interact with the Typewriter through its internal
// TypewriterController, which exposes the available actions.
package command

func Run() {
	typewriter := NewTypewriter()

	char := NewPressCharacter(typewriter)
	space := NewPressSpace(typewriter)
	tab := NewPressTab(typewriter, 4)
	enter := NewPressEnter(typewriter)
	backspace := NewPressBackspace(typewriter)

	// Invoke the commands with fluent interface:
	typewriter.
		Press(char.Key('H')).
		Press(char.Key('E')).
		Press(char.Key('R')).
		UndoLast().
		Press(char.Key('L')).
		Press(char.Key('L')).
		Press(char.Key('O')).
		Press(char.Key('!')).
		Press(enter).
		Press(enter).
		Press(tab).
		Press(char.Key('D')).
		Press(char.Key('E')).
		Press(space).
		Press(char.Key('R')).
		Press(backspace).
		Press(backspace).
		Press(char.Key('A')).
		Press(char.Key('R')).
		Press(enter).
		Press(enter).
		Press(char.Key(':')).
		Press(char.Key(')')).
		Press(backspace).
		UndoLast().
		Press(enter)

	// Print the result:
	typewriter.Print()
}
