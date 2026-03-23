package main

import "fmt"

type Jogada struct{
    pa, pb int;
}

/*
Em Go, na criação de uma função, o tipo vem depois da variável,
que nesse contexto é um struct. Isso serve para facilitar a
inserção dos tipos.
*/
func pontuacao(j Jogada) int{
    if(j.pa > j.pb){
        return j.pa - j.pb
    }

    return j.pb - j.pa
}

func main() {
    n := 0;
    fmt.Scan(&n);

    //Criando um vetor com parâmetro "n" deixando todas as posições com 0 
    jogadas := make([]Jogada, n)

    //Percorrendo o vetor para inserir os elementos via Scan
    for i := range jogadas{
        fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)
    }

    ind_melhor := -1;

    for i, jog := range jogadas{
        if jog.pa < 10 || jog.pb < 10{
            continue
        }

        //Ou ele é o primeiro ou ele é melhor do que o melhor
        if ind_melhor == -1 || pontuacao(jog) < pontuacao(jogadas[ind_melhor]){
            ind_melhor = i;
        }
    }

    if ind_melhor == -1{
        fmt.Println("sem ganhador")
    } else{
        fmt.Println(ind_melhor)
    }
}
