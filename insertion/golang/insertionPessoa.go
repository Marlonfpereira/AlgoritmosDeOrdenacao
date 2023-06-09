package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pessoa struct {
	codigo int
	name   string
}

func main() {
	input, err := os.Open("pessoa 1.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	output, err := os.OpenFile("./insGoPessoa.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	vetor := make([]Pessoa, 0)
	var comp, swap int

	num, err := strconv.Atoi(os.Args[1])

	scanner := bufio.NewScanner(input)
	for i := 0; i < num; i++ {

		scanner.Scan()
		linha := scanner.Text()
		valores := strings.Split(linha, "\t")
		id, _ := strconv.Atoi(valores[0])
		aux := Pessoa{id, valores[1]}
		vetor = append(vetor, aux)
	}

	// for _, v := range vetor {
	// 	fmt.Printf("%d %s, ", v.codigo, v.name)
	// }
	// println("\nsort\n")

	inicio := time.Now()
	for i := 0; i < len(vetor)-1; i++ {
		for j := i + 1; j > 0; j-- {
			comp++
			if vetor[j].codigo < vetor[j-1].codigo {
				var aux1 Pessoa
				aux1 = vetor[j]
				vetor[j] = vetor[j-1]
				vetor[j-1] = aux1
				swap++
			} else {
				break
			}
		}
	}
	fim := time.Now()

	// for _, v := range vetor {
	// 	fmt.Printf("%d %s\n", v.codigo, v.name)
	// }

	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)
}
