package main

import "fmt"

func main() {
	const maxCapacity = 1000
	var weights [maxCapacity]float64
	var n int

	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	if n <= 0 || n > maxCapacity {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan", maxCapacity)
		return
	}

	fmt.Println("Masukkan berat masing-masing anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	minWeight, maxWeight := weights[0], weights[0]
	for i := 1; i < n; i++ {
		if weights[i] < minWeight {
			minWeight = weights[i]
		}
		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", maxWeight)
}
