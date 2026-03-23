package main
import "fmt"

type Posicao struct{
    posX, posY int;
}

func main() {
    q := 0;
    c := "";
    fmt.Scan(&q, &c);

    //fmt.Println(q, c);

    posicoes := make([]Posicao, q);

    for i := range posicoes{
        fmt.Scan(&posicoes[i].posX, &posicoes[i].posY);
    }

    //fmt.Println(posicoes)
    
    anterior := posicoes[0];

    switch c {
        case "L":
            posicoes[0].posX--;
        case "R":
            posicoes[0].posX++;
        case "U":
            posicoes[0].posY--;
        case "D":
            posicoes[0].posY++;
    }

    for i := 1; i < q; i++{
        atual := posicoes[i];
        posicoes[i] = anterior
        anterior = atual;
    }

    for _, p := range posicoes{
        fmt.Println(p.posX, p.posY);
    }
    
}
