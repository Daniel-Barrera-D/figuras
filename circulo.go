package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (circulo *Circulo) calcularArea() float64 {
	return math.Pi * (circulo.Radio * circulo.Radio)
}

func (circulo *Circulo) calcularPerimetro() float64 {
	return (2 * math.Pi) * circulo.Radio
}
