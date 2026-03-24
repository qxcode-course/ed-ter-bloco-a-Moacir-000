package main
import "fmt"

func procurar_vivo(elementos []int, size, pos int){
    fmt.Println(size);
}

func main() {
    n, e := 0, 0;
    fmt.Scan(&n, &e);
    
    elementos := make([]int, n);

    procurar_vivo(elementos, n, e);
}
