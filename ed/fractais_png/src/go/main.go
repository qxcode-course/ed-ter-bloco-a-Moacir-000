package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func circulo(pen *Pen, raio float64){
	if raio < 10{
		return;
	}

	pen.Up();
	pen.Walk(150);
	pen.Down();

	pen.DrawCircle(raio / 2);

	pen.Up();
	pen.Walk(-150);
	pen.Down()

	pen.Right(60);
	circulo(pen, raio / 2);

}

func main() {
	pen := NewPen(500, 500);
	pen.SetRGB(0, 0, 0);
	pen.SetHeading(90);
	pen.SetPosition(250, 250);

	pen.DrawCircle(150);

	circulo(pen, 100);

	pen.SavePNG("circulos.png")
	fmt.Println("PNG file created successfully.")
}