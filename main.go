package main

import "task/task2"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
//  <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	/*//只出现一次的数字
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

	//数组操作、进位处理
	resT5 := taks1.T5([]int{9, 9, 9, 9})
	fmt.Println(resT5)

	//删除有序数组中的重复项
	resT6 := taks1.T6([]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 9})
	fmt.Println(resT6)

	//合并区间
	resT7 := taks1.T7([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(resT7)

	//两数之和
	resT8 := taks1.T8([]int{2, 7, 11, 15}, 9)
	fmt.Println(resT8)*/

	//-------------------------------------
	//指针 1
	/*num := 5
	task2.PlusTen(&num)
	fmt.Println(num)

	//指针 2
	point2 := []int{1, 2, 3, 4, 6}
	task2.MultiplyByTwo(&point2)
	fmt.Println(point2)

	//Goroutine 1
	task2.PrintOddEven()

	//Goroutine 2
	task2.TaskScheduler()

	//面向对象 1
	task2.Struct1()

	//面向对象 2
	task2.Struct2()

	//Channel 1
	task2.Channel1()
	//Channel 2
	task2.Channel2()*/
	task2.Lock1()
	task2.Lock2()
}
