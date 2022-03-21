package main

import "fmt"

/**
 * Created by : GoLand
 * User: ruohuai
 * Date: 2022/3/21
 * Time: 15:29
 */

//难度: 简单

//核心思想: 利用哈希表查询优势

func main() {
	fmt.Println(twoSum([]int{1, 3, 2, 7, 5}, 6))
}

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	var contains = make(map[int]int, numsLen)
	for i := 0; i < numsLen; i++ {
		if v, ok := contains[target-nums[i]]; ok { //加载数组第二个元素时,第一个元素已进map
			return []int{v, i}
		}
		contains[nums[i]] = i
	}
	return nil

}
