package figuras

type Cuadrado struct {
	Lado  float64
}


func (cuadrado *Cuadrado)calcularArea() float64 {
	return cuadrado.Lado * cuadrado.Lado
}

func (cuadrado *Cuadrado)calcularPerimetro() float64 {
	return 4 * cuadrado.Lado
}