package main

import (
	"fmt"
)

func main() {

	s := fmt.Sprintf("%.2f", 12.3456) // s == "12.35"
	fmt.Println(s)

	item := "Chocolate"
	price := 15.3333

	price_rounded := fmt.Sprintf("%.2f", price)
	item_price := fmt.Sprintf("%s%s%s%s%s", item, " ", "(", price_rounded, ")")
	print(item_price)
}
