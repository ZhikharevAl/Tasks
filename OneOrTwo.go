/*
Возвращать из функции вам нужно два значения. Если строка равна one, нужно вернуть первое число и саму строку.
Если строка равна two, нужно вернуть второе число и саму строку. Если строка другая - нужно вернуть 0 и саму строку.
*/

package main

func one_or_two(a int, b int, c string) (int, string) {
	var d int = 0
	var e string = ""
	switch c {
	case "one":
		d = a
		e = c
	case "two":
		d = b
		e = c
	default:
		e = c
	}
	return d, e
}
