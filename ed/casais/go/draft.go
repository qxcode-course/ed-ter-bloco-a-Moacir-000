package main
import "fmt"
func main() {
    n := 0;
    fmt.Scan(&n);

    animais := make([]int, n);
    //Declarando um map (chave e conteúdo)
    solteiros := make(map[int]int); 

    /*Inserção
    solteiros[-5] = 3
    */

    //Devolve o valor da chave e, caso não tenha, devolve 0 (default)
    //valor1 := solteiros[-3];
    //valor2 := solteiros[6];

    for i := range animais{
        fmt.Scan(&animais[i])
    }

    n_pares := 0;

    //O "_" identifica que não preciso desse valor
    for _, animal := range animais{
        if (solteiros[-animal] > 0){
            solteiros[-animal]--;
            n_pares++;
        } else{
            solteiros[animal]++;
        }
    }

    fmt.Println(n_pares);
}
