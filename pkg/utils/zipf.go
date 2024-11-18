package utils

import (
	"math"
	"math/rand"
)

// GenerateZipfDistribution generates subreddit popularity using Zipf distribution
func GenerateZipfDistribution(n int, s float64) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = int(math.Pow(float64(i+1), -s) * float64(rand.Intn(100)+1))
	}
	return result
}
