package main
import "fmt"

func resto(n int){
    if(n == 0){
        return;
    }

    r := n / 2;
    mod := n % 2;
    resto(r)
    fmt.Println(r, mod);
}

func main() {
    n := 0;
    fmt.Scan(&n);

    resto(n);
}
