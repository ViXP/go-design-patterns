package command

// TypewriterController implements all possible actions of the Typewriter receiver.
type TypewriterController struct {
	writer *Typewriter
}

// PressCharacter implements the behaviour logic for the simple character key press.
func (tc *TypewriterController) PressCharacter(char rune) {
	tc.writer.Draft += string(char)
}

// PressSpace implements the behaviour logic for the spacebar key press.
func (tc *TypewriterController) PressSpace() {
	tc.writer.Draft += " "
}

// PressTab implements the behaviour logic for the tab key press.
func (tc *TypewriterController) PressTab(tabSize byte) {
	for range tabSize {
		tc.writer.Draft += " "
	}
}

// PressBackspace implements the behaviour logic for the backspace key press.
func (tc *TypewriterController) PressBackspace() rune {
	draftLen := len(tc.writer.Draft)

	if draftLen == 0 {
		return 0
	}

	var removed rune = []rune(tc.writer.Draft[draftLen-1:])[0]
	tc.writer.Draft = tc.writer.Draft[:draftLen-1]
	return removed
}

// PressEnter implements the behaviour logic for the Enter key press.
func (tc *TypewriterController) PressEnter() {
	tc.writer.Draft += "\n"
}

// NewTypewriterController constructs a new TypewriterController instance.
func NewTypewriterController(tw *Typewriter) *TypewriterController {
	return &TypewriterController{tw}
}
