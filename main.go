package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10_000_000 // размер массива
	CHUNKS = 8          // количество частей для деления
)

// generateRandomElements генерирует слайс заданного размера.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = rand.Intn(100_000_000) // Случайные положительные числа
	}
	return res
}

// maximum находит максимальный элемент в слайсе.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks находит максимум, разбивая слайс на CHUNKS частей и используя горутины.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// Если данных меньше, чем частей, работаем в один поток
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		
		// Последний кусок забирает остаток элементов
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(chunk []int, idx int) {
			defer wg.Done()
			maxValues[idx] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Println("Старт программы...")
	data := generateRandomElements(SIZE)
	if len(data) == 0 {
		fmt.Println("Слайс пуст")
		return
	}

	// 1. Поиск в один поток
	start := time.Now()
	max1 := maximum(data)
	duration1 := time.Since(start).Microseconds()
	fmt.Printf("Однопоточный поиск: максимум = %d, время = %d мкс\n", max1, duration1)

	// 2. Многопоточный поиск
	start = time.Now()
	max2 := maxChunks(data)
	duration2 := time.Since(start).Microseconds()
	fmt.Printf("Многопоточный поиск: максимум = %d, время = %d мкс\n", max2, duration2)
}