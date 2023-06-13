package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


func QuickSort(values []int, began, end int, comp, swap *int) {
	i := began
	j := end -1
	pivo := values[(began + end) / 2]

	for i <= j {
		for values[i] < pivo && i < end {
			i++
			*comp+=2
		}
		for values[j] > pivo && j > began {
			j--
			*comp+=2
		}
		if i<=j {
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
}


func main() {
	input, err := os.Open("numeros 1.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	output, err := os.OpenFile("./quickGoNums.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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
	QuickSort(vetor, 0, num, &comp, &swap)
	fim := time.Now()


	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)

}
