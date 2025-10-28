package template_method

import "fmt"

// Radio represents a concrete device that implements the DeviceOperator interface, allowing it to be serviced
// through the maintenance process.
type Radio struct {
	owner        string
	in_order     bool
	currentCheck byte
	complexity   byte
}

// Assemble assembles the Radio after maintenance.
func (r *Radio) Assemble() {
	fmt.Printf("The beautiful radio of %s is assembled!\n", r.owner)
}

// CheckElectronics inspects and repairs the internal electronics of the Radio.
func (r *Radio) CheckElectronics() {
	fmt.Println("Checking the radio's electronics...")
	if r.complexity > r.currentCheck {
		fmt.Printf(
			"The electronics of the Radio is fixed for the %vth time, some details are changed\n", r.currentCheck+1,
		)
		r.currentCheck++
	} else {
		r.in_order = true
		fmt.Println("All inner electronics is in working condition!")
	}
}

// Clean cleans the Radio's interior and components.
func (r *Radio) Clean() {
	fmt.Println("The radio is cleaned with compressed air spray.")
}

// Disassemble disassembles the Radio for maintenance.
func (r *Radio) Disassemble() {
	fmt.Printf("The beautiful radio of %s is disassembled!\n", r.owner)
}

// Run tests whether the Radio is operational.
func (r *Radio) Run() bool {
	if r.complexity == 0 {
		r.in_order = true
	}

	if r.in_order {
		fmt.Println("The radio works fine!")
	} else {
		fmt.Println("The radio is still not working :(")
	}
	return r.in_order
}

// NewRadio creates the new Radio with the specific owner and maintenance complexity.
func NewRadio(owner string, complexity byte) *Radio {
	return &Radio{owner, false, 0, complexity}
}

// Interface implementation assertion.
var _ DeviceOperator = &Radio{}
