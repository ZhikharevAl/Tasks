package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	input := []int{a, b, c}
	text := [11]string{"Ноль", "Один", "Два", "Три", "Четыре", "Пять", "Шесть", "Семь", "Восемь", "Девять", "Десять"}
	for _, v := range input {
		for i, _ := range text {
			if v == i {
				fmt.Println(text[i])
			}
		}
	}
}
