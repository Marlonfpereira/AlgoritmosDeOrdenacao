package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	input, err := os.Open("nomes 1.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	output, err := os.OpenFile("./insGoNomes.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	vetor := make([]string, 0)
	var comp, swap int

	scanner := bufio.NewScanner(input)
	num, err := strconv.Atoi(os.Args[1])
	for i := 0; i < num; i++ {
		var aux string
		scanner.Scan()
		aux = scanner.Text()
		vetor = append(vetor, aux)
	}

	// for _, v := range vetor {
	// 	fmt.Printf("%s, ", v)
	// }
	// println("\nsort\n")

	inicio := time.Now()
	for i := 0; i < len(vetor)-1; i++ {
		for j := i + 1; j > 0; j-- {
			comp++
			if vetor[j] < vetor[j-1] {
				var aux1 string
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
	// 	fmt.Printf("%s, ", v)
	// }

	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)
}
