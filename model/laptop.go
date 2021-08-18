package model

type Laptop struct {
	brand string
	model string
	cpu   string
	ram   string
}

/* Constructor BEGIN */
func (l *Laptop) Create(data []string) *Laptop {
	return &Laptop{
		brand: data[0],
		model: data[1],
		cpu:   data[2],
		ram:   data[3],
	}
}

/* Brand Getter */
func (l Laptop) Brand() string {
	return l.brand
}

/* Model Getter */
func (l Laptop) Model() string {
	return l.model
}

/* CPU Getter */
func (l Laptop) Cpu() string {
	return l.cpu
}

/* RAM Getter */
func (l Laptop) Ram() string {
	return l.ram
}
