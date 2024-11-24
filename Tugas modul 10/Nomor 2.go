package main

import "fmt"

func main() {
	const maxCapacity = 1000
	var weights [maxCapacity]float64
	var x, y int

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > maxCapacity {
		fmt.Println("Jumlah ikan (x) harus antara 1 dan", maxCapacity)
		return
	}
	if y <= 0 {
		fmt.Println("Kapasitas wadah (y) harus lebih besar dari 0")
		return
	}

	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	totalContainers := (x + y - 1) / y
	containerWeights := make([]float64, totalContainers)
	for i := 0; i < x; i++ {
		containerIndex := i / y
		containerWeights[containerIndex] += weights[i]
	}

	var totalWeight float64
	for _, weight := range containerWeights {
		totalWeight += weight
	}
	averageWeight := totalWeight / float64(totalContainers)

	fmt.Println("Total berat ikan di setiap wadah:")
	for i, weight := range containerWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, weight)
	}

	fmt.Printf("Berat rata-rata per wadah: %.2f\n", averageWeight)
}
