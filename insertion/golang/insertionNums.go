package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	input, err := os.Open("numeros 1.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	output, err := os.OpenFile("./insGoNums.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	
	vetor := make([]int, 0)

	var comp, swap int

	num, err := strconv.Atoi(os.Args[1])

	for i := 0; i < num; i++ {
		var aux int
		fmt.Fscan(input, &aux)
		vetor = append(vetor, aux)
	}

	inicio := time.Now()
	
	for i := 0; i < len(vetor)-1; i++ {
		for j := i + 1; j > 0; j-- {
			comp++
			if vetor[j] < vetor[j-1] {
				var aux1 int
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

	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)
}
