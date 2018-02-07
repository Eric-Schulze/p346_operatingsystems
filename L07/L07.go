package L07

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func SortStrings(args []string) (words []string, nums []float64) {
	for i := 0; i < len(args); i++ {
		if n, err := strconv.ParseFloat(args[i], 64); err == nil {
			nums = append(nums, n)
		} else {
			words = append(words, args[i])
		}
	}

	if words == nil {
		words = []string{}
	}
	if nums == nil {
		nums = []float64{}
	}
	return words, nums
}

func SumSlice(nums []float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	//Retrieve command line arguments
	args := os.Args[1:]

	//Sort arguments into strings and numbers
	words, nums := SortStrings(args)

	//Sort words slice
	sort.Strings(words)

	//Sum nums slice
	sum := SumSlice(nums)

	//Print Outputs
	fmt.Printf("Input list: %s\n\n", args)
	fmt.Printf("Sorted Words:\n")
	for _, s := range words {
		fmt.Printf("%s\n", s)
	}
	fmt.Println()
	fmt.Printf("The sum of the numbers is %f\n\n", sum)
}
