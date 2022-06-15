package main

import (
	"fmt"
)

func main() {
	appleCost := 5.99 // Одне яблуко коштує 5.99 грн
	pearCost := 7.    // Ціна однієї груші - 7 грн
	myBalance := 23.  // Ми маємо 23 грн
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println(float64(9*appleCost+8*pearCost), "грн")
	fmt.Println("2. Скільки груш ми можемо купити?")
	fmt.Println(int(myBalance/pearCost), "шт.")
	fmt.Println("3.  Скільки яблук ми можемо купити?")
	fmt.Println(int(myBalance/appleCost), "шт.")
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	fmt.Println(2*appleCost+2*pearCost <= myBalance)
}
