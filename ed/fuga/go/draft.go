package main

import "fmt"

func main() {
	H, P, F, D := 0, 0, 0, 0

	fmt.Scan(&H, &P, &F, &D)

	for i := 0; i < 25; i++ {
		F = F + D
		if F == H {
			fmt.Println("S")
			break
		} else if F == P {
			fmt.Println("N")
			break
		} else if(F == 0 && D == -1){
            F = 16;
        } else if(F == 15 && D == 1){
            F = -1;
        }
	}
}
