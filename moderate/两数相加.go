package main

import "fmt"

/**
 * Created by : GoLand
 * User: ruohuai
 * Date: 2022/3/21
 * Time: 15:43
 */

//难度: 中等

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Append(val int) *ListNode {
	l.Next = &ListNode{Val: val}
	return l.Next
}

func main() {
	var l1 = &ListNode{Val: 1}
	l1.Append(2).Append(3)
	//此时l1代表的数字为 321
	var l2 = &ListNode{Val: 5}
	l2.Append(6).Append(7)
	//此时l2代表的数字为 765
	l3 := addTwoNumbers(l1, l2)
	var sum string
	for l3 != nil {
		sum += fmt.Sprint(l3.Val)
		l3 = l3.Next
	}

	fmt.Println("The sum of l1 plus l2 is", reversalString(sum))
}

func reversalString(input string) string {
	inputLength := len(input)
	var bytes = make([]byte, 0, inputLength)
	for i := inputLength - 1; i >= 0; i-- {
		bytes = append(bytes, input[i])
	}
	return string(bytes)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList1, newList2 *ListNode
	var carryVal int

	for l1 != nil || l2 != nil {
		var l1ElemVal, l2ElemVal = 0, 0
		if l1 != nil {
			l1ElemVal = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2ElemVal = l2.Val
			l2 = l2.Next
		}
		sum := l1ElemVal + l2ElemVal + carryVal
		sum, carryVal = sum%10, sum/10
		if newList1 == nil {
			newList1 = &ListNode{Val: sum}
			newList2 = newList1

		} else {
			newList2.Next = &ListNode{Val: sum}
			newList2 = newList2.Next
		}
	}
	if carryVal > 0 {
		newList2.Next = &ListNode{
			Val: carryVal,
		}
	}
	return newList1
}
