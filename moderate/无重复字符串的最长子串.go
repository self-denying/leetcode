package main

import "fmt"

/**
 * Created by : GoLand
 * User: ruohuai
 * Date: 2022/3/21
 * Time: 17:44
 */

func main() {
	fmt.Println(lengthOfLongestSubstring("helabcab"))
}

func lengthOfLongestSubstring(s string) int {
	sl := len(s)
	var cursor = 0  //定义游标
	var counter int //定义计数器
	var set = make(map[byte]int)
	for i := 0; i < sl; i++ {
		if i != 0 {
			delete(set, s[i-1]) //除了第一次,因为上一个元素已匹配完毕,此时应该匹配与下一个元素不相同的其他字符
		}
		for cursor < sl && set[s[cursor]] == 0 {
			set[s[cursor]]++
			cursor++
		}
		counter = max(counter, cursor-i) //cursor-i表示两个字符之间的距离
	}

	return counter

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
