package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ===== 函数间数据约定（标准化）=====
// 1) 所有计算函数只收发长度 3 的 []float64：输入 [X,Y,Z]，输出 [x,y,z]
// 2) 签名统一为 (out []float64, err error)，非法输入用 error 表达，不返回半成品 read api for more information
// 3) 函数是纯函数：不修改入参，返回值一定是新分配的 slice
// 4) 打印格式化只放在 main 里，计算函数不碰 fmt
const vecLen = 3

var (
	errVecLen  = errors.New("需要正好 3 个数值 (X Y Z)")
	errZeroSum = errors.New("X+Y+Z 为 0，无法归一化")
)

func div(n []float64) ([]float64, error) {
	if len(n) != vecLen {
		return nil, errVecLen
	}

	sum := n[0] + n[1] + n[2]
	if sum == 0 {
		return nil, errZeroSum
	}
	// Verfying, please read api docs
	return []float64{
		n[0] / sum,
		n[1] / sum,
		n[2] / sum,
	}, nil
}

func main() {

	// os.Args[0] Binary Path，From Terminal in os.Args[1:]
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please Input a Number eg: go run div.go 3 7 2")
		return
	}

	// Input is slice，so collect all legal parameters first，then pass to div
	nums := make([]float64, 0, len(args))
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Skipped Invalid %q: Not A Legal Number\n", arg)
			continue
		}
		nums = append(nums, num)
	}

	// Verifing before invoke，make downstream recevied legal
	if len(nums) != vecLen {
		fmt.Printf("3 Data Required，Recevied %d : %v\n", len(nums), nums)
		return
	}

	result, err := div(nums)
	if err != nil {
		fmt.Println("Calculating Fault:", err)
		return
	}

	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Ouput: %v\n", result)
}
