package template_method

// DeviceOperator declares the common interface for the devices that could be maintained.
type DeviceOperator interface {
	Disassemble()
	Assemble()
	CheckElectronics()
	Clean()
	Run() bool
}

// Maintainer defines the Template Method for performing maintenance on a device.
type Maintainer interface {
	Maintain(d DeviceOperator)
}
