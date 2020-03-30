package main

import (
	"fmt"
	"sort"
)

//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//
//示例 1:
//输入: [1,2,0]
//输出: 3

//示例 2:
//输入: [3,4,-1,1]
//输出: 2

//示例 3:
//输入: [7,8,9,11,12]
//输出: 1
//

func main() {
	//nums := []int{-1}
	nums := []int{0, 2, 2, 1, 1}
	//nums := []int{1, 2, 0}
	//nums := []int{3,4,-1,1}
	//nums := []int{7,8,9,11,12}
	positive := firstMissingPositive(nums)
	fmt.Println(positive)
}

func firstMissingPositive(nums []int) int {

	sort.Ints(nums)
	var tmp []int
	//去重 去负数
	for i := range nums {
		flag := true
		for j := range tmp {
			if nums[i] == tmp[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag && nums[i] > 0 { // 标识为false，不添加进结果
			tmp = append(tmp, nums[i])
		}
	}

	if len(tmp) == 0 {
		return 1
	}

	for i := 0; i < len(tmp); i++ {
		if i+1 < tmp[i] {
			return i + 1
		} else {
			if i+1 == tmp[i] {
				if tmp[i] == tmp[len(tmp)-1] {
					return tmp[i] + 1
				}
				if tmp[i] == tmp[i+1]-1 {
					continue
				} else {
					return tmp[i] + 1
				}
			}
		}
	}
	return 0
}
