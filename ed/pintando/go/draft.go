package main
import "fmt"
import "math"
func main() {
    a, b, c := 0.0, 0.0, 0.0;
    fmt.Scan(&a, &b, &c);

    p := (a + b + c) / 2;

    area := math.Sqrt(p * (p - a) * (p - b) * (p - c));

    fmt.Printf("%.2f\n", area);
}
