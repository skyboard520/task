package main

import (
	"fmt"
	"task/taks1"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
//  <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//只出现一次的数字
	var nums []int = []int{1, 2, 3, 3, 4, 5, 4}
	result := taks1.T1(nums)
	fmt.Println(result)

	//回文数
	b := taks1.T2("我是我")
	fmt.Println(b)

	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	b = taks1.T3("我我")
	fmt.Println(b)

	//最长公共前缀
	stirs := []string{"flower", "flow", "flight"}
	res := taks1.T4(stirs)
	fmt.Println(res)

}
