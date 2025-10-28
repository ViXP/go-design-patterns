package template_method

import "fmt"

// MaintenanceTechnician is a concrete Maintainer that performs maintenance using the Maintain Template Method.
type MaintenanceTechnician struct{}

// Maintain defines the template method for maintaining a device. It outlines the general algorithm for maintenance,
// while the specific behavior of each step depends on the concrete DeviceOperator implementation.
func (m *MaintenanceTechnician) Maintain(d DeviceOperator) {
	d.Disassemble()
	d.Clean()

	for !d.Run() {
		d.CheckElectronics()
	}

	d.Assemble()
	fmt.Println("---")
}

// Interface implementation assertion.
var _ Maintainer = &MaintenanceTechnician{}
