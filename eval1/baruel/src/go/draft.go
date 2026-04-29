package main
import "fmt"

func main() {

    //OBS: acabei esquecendo a formatação de saída e como expor os faltantes.
    var qtd_total, qtd_possui int;

    fmt.Scan(&qtd_total, &qtd_possui);

    figuras := make([]int, qtd_possui);

    figuras_repetidas := make([]int, 0);

    for i := range figuras{
        fmt.Scan(&figuras[i]);
    }

    map_repetidas := make(map[int]bool);

    n_repetidas := 0;

    for _, figura := range figuras{
        if(map_repetidas[figura]){
            figuras_repetidas = append(figuras_repetidas, figura)
            n_repetidas++;
        } else{
            map_repetidas[figura] = true;
        }
    }

    figuras_faltantes := make([]int, 0);

    //map_faltantes := make(map[int]bool);

    for _, valor := range figuras{
        if(map_repetidas[valor]){
            figuras_faltantes = append(figuras_faltantes, valor);
        } else{
            map_repetidas[valor] = false;
        }
    }

    if(len(figuras_repetidas)) == 0{
        fmt.Println("N");
    } else{
        fmt.Println(figuras_repetidas);
        //fmt.Println(figuras_faltantes);
        fmt.Println(n_repetidas);
    }

    fmt.Println(figuras_faltantes);
}
