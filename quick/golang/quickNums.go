package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func partition(arr []int, low, high int, comp, swap *int) ([]int, int) {
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

func quickSort(arr []int, low, high int, comp, swap *int) []int {
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

	// for _, v := range vetor {
	// 	print(v)
	// 	print(" ")
	// }
	// println()

	inicio := time.Now()
	quickSort(vetor, 0, num-1, &comp, &swap)
	fim := time.Now()

	// for _, v := range vetor {
	// 	print(v)
	// 	print(" ")
	// }

	fmt.Fprintf(output, "%s,%v,%d,%d\n", os.Args[1], float64(fim.Sub(inicio).Seconds()), comp, swap)

}
