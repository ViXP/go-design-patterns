// Package template_method implements the Template Method design pattern.
// The Maintainer interface defines the Template Method, implemented by the concrete MaintenanceTechnician type. This
// type defines the abstract sequence of steps for maintaining various devices.
// All devices implement the DeviceOperator interface, which declares the set of operations expected by the Template
// Method. This allows polymorphic execution of the maintenance process across different device types (e.g., Radio,
// Television) while keeping the overall algorithm structure fixed.
package template_method

func Run() {
	// Create multiple devices in different shape:
	brokenTv := NewTelevision("Dewey's TV", false)
	workingTv := NewTelevision("Gale's TV", true)
	problematicRadio := NewRadio("Sidney", 5)
	fineRadio := NewRadio("Randy", 0)

	// Create maintainance worker:
	worker := &MaintenanceTechnician{}

	// Maintain the devices one by one:
	worker.Maintain(brokenTv)
	worker.Maintain(workingTv)
	worker.Maintain(problematicRadio)
	worker.Maintain(fineRadio)
}
