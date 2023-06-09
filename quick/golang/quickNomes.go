package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func partition(arr []string, low, high int, comp, swap *int) ([]string, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		*comp++
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			*swap++
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	*swap++
	return arr, i
}

func quickSort(arr []string, low, high int, comp, swap *int) []string {
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
	input, err := os.Open("nomes 1.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	output, err := os.OpenFile("./quickGoNomes.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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

	quickSort(vetor, 0, num-1, &comp, &swap)

	fim := time.Now()

	// for _, v := range vetor {
	// 	fmt.Printf("%s, ", v)
	// }

	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)
}
