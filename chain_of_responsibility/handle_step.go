package chain_of_responsibility

// handleStep is the helper function that allows all cooking steps to share the handling logic with correct process
// calling.
func handleStep(step Chainable) bool {
	step.process()
	if step.getNext() != nil {
		return step.getNext().Handle()
	}
	return true
}
