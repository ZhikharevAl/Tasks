/*
Напишите функцию double_m(), которая должна принимать на вход два целых числа a и b
и возвращать сумму квадратов чисел от a до b (включительно).
*/

package main

func double_m(a, b int) int {
	var sum int = 0
	if a < b {
		for i := a; i <= b; i++ {
			sum += i * i
		}
	} else {
		for i := b; i <= a; i++ {
			sum += i * i
		}
	}
	return sum
}
