/*
Ниже приведен код Go, который вычисляет общую цену с учетом налога и определяет,
хватит ли имеющихся денег для покупки.
*/

package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = taxRate * 100
	fmt.Println("Tax is", int(tax), "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", int(total) < availableFunds)
}
