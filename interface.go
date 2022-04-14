package figuras

import "fmt"

type Figura interface {
	calcularArea() float64
	calcularPerimetro() float64
}

func CalculoFigura(figura Figura) {
	fmt.Println("Medidas:",figura)
	fmt.Println("Área:",figura.calcularArea())
	fmt.Println("Perímetro:",figura.calcularPerimetro())
}