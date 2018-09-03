package leetcode

// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
// 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

// 示例：给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

// 1.暴力法
// 遍历每个元素x
// 复杂度：时间复杂度：O(n^2)
// 执行用时：60 ms
func TwoSum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 2.一次遍历法
// 时间复杂度：O(n)
// 执行用时：100 ms
func TwoSum2(nums []int, target int) []int {
	length := len(nums)
	j := 0
	for i := j + 1; i < length; {
		if target == nums[j] + nums[i] {
			return []int{j, i}
		}
		i++
		if i >= length {
			j++
			i = j + 1
		}
	}
	return []int{}
}

// 3.一遍map表
// 把遍历过的元素存进map的key中，然后再判断map中是否存在
// 执行用时：8 ms
func TwoSum3(nums []int, target int) []int {
	length := len(nums)
	m := make(map[int]int)
	for i := 0; i < length; i++ {
		key := target - nums[i]
		_, ok := m[key]
		if ok {
			return []int{m[key], i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
