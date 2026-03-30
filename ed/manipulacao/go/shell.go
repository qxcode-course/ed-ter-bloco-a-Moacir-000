package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int{
	if n < 0{
		return -n
	}
	return n
}

func getMen(vet []int) []int {
	final := make([]int, 0)
	for i := range vet{
		if (vet[i] > 0){
			final = append(final, vet[i])
		}
	}

	return final
}

func getCalmWomen(vet []int) []int {
	final := make([]int, 0);

	for i := range vet{
		if (vet[i] > -10 && vet[i] < 0){
			final = append(final, vet[i]);
		}
	}
	return final
}

func sortVet(vet []int) []int {
	n := len(vet)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if vet[j] > vet[j+1] {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}

	return vet
}

func sortStress(vet []int) []int {
	n := len(vet)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if abs(vet[j]) > abs(vet[j+1]) {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}

	return vet
}

func reverse(vet []int) []int {
	final := make([]int, 0);

	tamanho := len(vet)
	tamanho = tamanho - 1;
	
	for i := tamanho; i >= 0; i--{
		final = append(final, vet[i])
	}

	return final
}

func unique(vet []int) []int {
	final := make([]int, 0);

	unicos := make(map[int]bool)

	for _, v := range vet{
		if !unicos[v]{
			final = append(final, v)
			unicos[v] = true;
		}
	}
	return final
}

func repeated(vet []int) []int {
	final := make([]int, 0);

	repetidos := make(map[int]int);

	for _, v := range vet{
		repetidos[v]++;

		if repetidos[v] > 1{
			final = append(final, v)
		}
	}
	return final
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

