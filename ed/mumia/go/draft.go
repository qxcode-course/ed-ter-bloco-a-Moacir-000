package main
import "fmt"
func main() {
    nome := "";
    idade := 0;

    fmt.Scan(&nome);
    fmt.Scan(&idade);

    if(idade < 12){
        fmt.Println(nome + " eh crianca");
    } else if(idade >= 12 && idade < 18){
        fmt.Println(nome + " eh jovem");
    } else if(idade >= 18 && idade < 65){
        fmt.Println(nome + " eh adulto");
    } else if(idade >= 65 && idade < 1000){
        fmt.Println(nome + " eh idoso");
    } else{
        fmt.Println(nome + " eh mumia");
    }
}
