package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	//var nums = []int{25, 2, 22, 34, 24, 30, 0, 14, 19, 19, 21, 3, 4, 2}
	//var ops = []string{"^", "-", ">>", "*", "**"}
	var nums = []int{19, 23, 2, 21, 31, 29, 2, 4, 965, 21, 2, 33, 7, 32, 28, 29, 19, 2, 22, 5, 34, 24, 30, 0, 19, 3, 4, 2}
	var ops = []string{"^", "^", "|", "/", "-", "%", "&", ">>", "*"}
	//var ops = []string{"^", "^", "|", "/", "-", "%", "&", ">>"}
	for _, s := range Is1024(nums, ops) {
		fmt.Println(s)
	}
}

func Is1024(nums []int, ops []string) []string { //1)检索的空间为4+3。还是挺大的，但是可以暴力解。要快的话，可以双向解，从1024倒着来。这样可以降低驱动的数据集。
	var hash = make(map[string]int)
	var set = make(map[string]int)
	var n, m = len(nums), len(ops)
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			if a == b {
				continue
			}
			for c := 0; c < n; c++ {
				if a == c || b == c {
					continue
				}
				for d := 0; d < n; d++ {
					if a == d || b == d || c == d {
						continue
					}
					for i := 0; i < m; i++ {
						for j := 0; j < m; j++ {
							for k := 0; k < m; k++ {
								if i == j || i == k || j == k {
									continue
								}
								mid, err := Calc(nums[a], nums[b], ops[i])
								if err != nil {
									//panic(err)
									continue
								}
								then, err := Calc(mid, nums[c], ops[j])
								if err != nil {
									//panic(err)
									continue
								}
								final, err := Calc(then, nums[d], ops[k])
								if err != nil {
									//panic(err)
									continue
								}
								if final == 1024 {
									set[fmt.Sprintf("%d, %s, %d, %s, %d, %s, %d", nums[a], ops[i], nums[b], ops[j], nums[c], ops[k], nums[d])] = 1
									hash[strconv.Itoa(nums[a])]++
									hash[strconv.Itoa(nums[b])]++
									hash[strconv.Itoa(nums[c])]++
									hash[strconv.Itoa(nums[d])]++
									hash[ops[i]]++
									hash[ops[j]]++
									hash[ops[k]]++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(hash)
	var ret = make([]string, len(set))
	for k := range set {
		ret = append(ret, k)
	}
	return ret
}

func Calc(a, b int, op string) (int, error) {
	switch op {
	case "^":
		return a ^ b, nil
	case "&":
		return a & b, nil
	case "|":
		return a | b, nil
	case "-":
		return a - b, nil
	case "+":
		return a + b, nil
	case "*":
		return a * b, nil
	case "%":
		if b == 0 {
			return 0, errors.New("0 divided")
		}
		return a % b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("0 divided")
		}
		return a / b, nil
	case ">>":
		return a >> b, nil
	case "<<":
		return a << b, nil
	case "**":
		p := math.Pow(float64(a), float64(b))
		if p >= 1<<31 {
			return 0, fmt.Errorf("overflow int32: %d ** %d = %f", a, b, p)
		}
		return int(p), nil
	default:
		panic("unknown op")
		//return 0, errors.New("unknown op")
	}
}
