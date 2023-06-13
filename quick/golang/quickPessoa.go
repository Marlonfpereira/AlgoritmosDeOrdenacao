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

func QuickSort(values []Pessoa, began, end int, comp, swap *uint64) []Pessoa {
	i := began
	j := end - 1
	pivo := values[(began+end)/2]

	for i <= j {
		for values[i].codigo < pivo.codigo || (values[i].codigo == pivo.codigo && values[i].name < pivo.name) {
			i++
			*comp += 3
		}
		for values[j].codigo > pivo.codigo || (values[j].codigo == pivo.codigo && values[j].name > pivo.name) {
			j--
			*comp += 3
		}
		if i <= j {
			values[i], values[j] = values[j], values[i]
			*swap++
			i++
			j--
		}
	}

	if j > began {
		QuickSort(values, began, j+1, comp, swap)
	}
	if i < end {
		QuickSort(values, i, end, comp, swap)
	}
	return values
}

func main() {
	input, err := os.Open("pessoa 1.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	output, err := os.OpenFile("./quickGoPessoa.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	vetor := make([]Pessoa, 0)
	var comp, swap uint64

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


	inicio := time.Now()
	QuickSort(vetor, 0, num, &comp, &swap)
	fim := time.Now()


	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)
}
