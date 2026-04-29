package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func distribuir(pen *Pen, raio float64, passos int) {
	if passos == 0 {
		return
	}

	pen.Up()
	pen.Walk(raio)
	pen.Down()

	circulo(pen, raio/2)

	pen.Up()
	pen.Walk(-raio)
	pen.Down()

	pen.Right(60)

	distribuir(pen, raio, passos-1)
}

func circulo(pen *Pen, raio float64){
	if raio < 10 {
		return
	}

	// desenha o círculo atual
	pen.DrawCircle(raio)

	// distribui os próximos círculos ao redor
	distribuir(pen, raio, 6)
}

func main() {
	pen := NewPen(500, 500);
	pen.SetRGB(0, 0, 0);
	pen.SetHeading(90);
	pen.SetPosition(250, 250);

	//pen.DrawCircle(150);

	circulo(pen, 50);

	pen.SavePNG("circulos.png")
	fmt.Println("PNG file created successfully.")
}