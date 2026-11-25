package main

import (
	"fmt"
	"os"
	"strconv"
)

func sub(n []float64) []float64 {
	var tmp0 []float64
	tmp0[0] = n[0] / (n[0] + n[1] + n[2])
	tmp0[1] = n[1] / (n[0] + n[1] + n[2])
	tmp0[2] = n[2] / (n[0] + n[1] + n[2])

	return tmp0[]
}

func main() {

	// os.Args[0] Binary Path，From Terminal in os.Args[1:]
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please Input a Number eg: go run sub.go 3 7 2")
		return
	}

	// Input is slice，so collect all leagal parameters first，then to  sub
	nums := make([]float64, 0, len(args))
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Skipped Unvalid Numbers %q: Not A Legal Number\n", arg)
			continue
		}
		nums = append(nums, num)
	}

	if len(nums) == 0 {
		fmt.Println("No Avaliable Number")
		return
	}


	result := sub(nums)
	fmt.Printf("输入: %v\n", nums)
	fmt.Printf("输出: %v\n", result)
}

