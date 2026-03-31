package main
import "fmt"

func main() {
    n, e := 0, 0;
    fmt.Scan(&n, &e);
    
    elementos := make([]int, n);

    for i := range elementos{
        elementos[i] = i+1;
    }

    i := e - 1;

    for len(elementos) > 1{
        fmt.Print("[ ");
        for j, v := range elementos{
            if j == i{
                fmt.Printf("%d> ", v);
            } else {
                fmt.Printf("%d ", v);
            }
        }
    
        fmt.Println("]")

        alvo := (i + 1) % len(elementos)

        elementos = append(elementos[:alvo], elementos[alvo+1:]...)

        i = alvo % len(elementos);
    }

    fmt.Printf("[ %d> ]\n", elementos[0]);
}