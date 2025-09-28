package taks1

/*
*
只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func T1(nums []int) []int {
	var result []int
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

/*
- 回文数
考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func T2(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/**
- **有效的括号 **
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/

func T3(s string) bool {
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []rune
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) == 0 || m[v] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

/*
*
- 最长公共前缀
考察：字符串处理、循环嵌套
题目：查找字符串数组中的最长公共前缀
*/
func T4(stirs []string) string {
	if len(stirs) == 0 {
		return ""
	}

	var runes = []rune(stirs[0])
	for i := 1; i < len(stirs); i++ {
		str := []rune(stirs[i])
		for j := 0; j < len(runes) && j < len(str); j++ {
			if runes[j] != str[j] {
				runes = runes[:j]
				break
			}
		}
	}

	return string(runes)
}
