package template_method

import "fmt"

// Television represents a concrete device that implements the DeviceOperator interface, allowing it to be serviced
// through the maintenance process.
type Television struct {
	name     string
	in_order bool
}

// Assemble assembles the Television after maintenance.
func (t *Television) Assemble() {
	fmt.Printf("The %v is assembled successfully.\n", t.name)
}

// CheckElectronics inspects and repairs the internal electronics of the Television.
func (t *Television) CheckElectronics() {
	if !t.in_order {
		fmt.Println("The electronics of the TV is checked, some transistors are changed")
		t.in_order = true
	} else {
		fmt.Println("All inner electronics is in working condition.")
	}
}

// Clean cleans the Television's interior and components.
func (t *Television) Clean() {
	fmt.Println("The TV is cleaned with vacuum cleaner and brushes")
}

// Disassemble disassembles the Television for maintenance.
func (t *Television) Disassemble() {
	fmt.Printf("The %v is disassembled\n", t.name)
}

// Run tests whether the Television is operational.
func (t *Television) Run() bool {
	if t.in_order {
		fmt.Println("The TV works fine!")
	} else {
		fmt.Println("The TV is still not working :(")
	}
	return t.in_order
}

// NewTelevision creates the new Television with the specific name in predefined working state.
func NewTelevision(name string, workingStatus bool) *Television {
	return &Television{name, workingStatus}
}

// Interface implementation assertion.
var _ DeviceOperator = &Television{}
