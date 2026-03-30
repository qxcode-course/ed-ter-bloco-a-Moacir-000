package main
import "fmt"

func imprimir(lista []int) {
    if len(lista) == 0 {
        fmt.Println("N")
        return;
    }
    for _, v := range lista {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}

func main() {
	n := 0;
	fmt.Scan(&n)

	fila := make([]int, n);

	for i := range fila{
		fmt.Scan(&fila[i])
	}

	m := 0;
	fmt.Scan(&m);

	fila_saida := make([]int, m);

	for i := range fila_saida{
		fmt.Scan(&fila_saida[i]);
	}

	pessoa_saiu := make(map[int]bool)

	for _, v := range fila_saida{
		pessoa_saiu[v] = true;
	}

	fila_nova := make([]int, 0)

	for _, pessoa := range fila{
		if !pessoa_saiu[pessoa]{
			fila_nova = append(fila_nova, pessoa)
		}
	}

	imprimir(fila_nova)

	/*fmt.Println(n);
	fmt.Println(fila);
	fmt.Println(m);
	fmt.Println(fila_saida);*/
}
