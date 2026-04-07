package main
import "fmt"

/* Usando ponteiros
func contarDigitosRecursivo(num int, qtd *int){
    if(num == 0){ //ponto de parada
        return;
    }

    *qtd += 1; //ação
    contarDigitosRecursivo(num / 10, qtd) //chamada recursiva 

}
*/

func contarDigitosRecursivo(num int, qtd int) int{
    if(num == 0){ //ponto de parada
        return qtd;
    }
    return contarDigitosRecursivo(num / 10, qtd + 1) //chamada recursiva 

}

func contarDigitos(num int) int{
    qtd := 0;
    for{
        if(num == 0){ //ponto de parada
            break;
        }

        qtd += 1; //ação
        num = num / 10; //diminuição do problema
    }

    return qtd;
}

func main() {
    num := 4891236;

    fmt.Println(contarDigitosRecursivo(num, 0));
}
