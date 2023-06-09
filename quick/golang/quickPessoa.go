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

func partition(arr []Pessoa, low, high int, comp, swap *int) ([]Pessoa, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		*comp++
		if arr[j].codigo < pivot.codigo {
			arr[i], arr[j] = arr[j], arr[i]
			*swap++
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	*swap++
	return arr, i
}

func quickSort(arr []Pessoa, low, high int, comp, swap *int) []Pessoa {
	*comp++
	if low < high {
		var p int
		arr, p = partition(arr, low, high, comp, swap)
		arr = quickSort(arr, low, p-1, comp, swap)
		arr = quickSort(arr, p+1, high, comp, swap)
	}
	return arr
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
	// 	fmt.Printf("%d %s\n", v.codigo, v.name)
	// }
	// println()

	inicio := time.Now()
	quickSort(vetor, 0, num-1, &comp, &swap)
	fim := time.Now()

	// for _, v := range vetor {
	// 	fmt.Printf("%d %s\n", v.codigo, v.name)
	// }

	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)
}
