/*
Реализуйте функцию SafeWrite(nums [5]int, i, val int) [5]int, которая записывает значение val в массив nums
по индексу i, если индекс находится в рамках массива. В противном случае массив возвращается без изменения.
*/

package main

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i >= 0 && i < len(nums) {
		nums[i] = val
	}
	return nums
}
