package main

import (
	"fmt"
	"main/clients"
	"main/products"
	"main/students"
	"golang.org/x/tour/tree"
	"main/functions"
)

func main() {
	fmt.Println("Hello, World!")

	hiro := students.Create("Geeko", "Hiro", "Computer Science", 20)

	hiro.DisplayDetails()

	haojue := products.New("Haojue", "A spidy motorcycle", 700000.0)
	purse := clients.NewPurse("Gogo Purse", "Black", 1000000.0)
	client := clients.New("Gogo", "Japan", 20)

	client.BuyProduct(haojue)

	client.SetPurse(&purse)
	client.BuyProduct(haojue)
	fmt.Println("Amount purse", client.Purse.Amount)

	t1 := tree.New(5) // generate a tree with values k, 2k, .. 10k (k=5)
	t2 :=  tree.New(2)

	fmt.Println(functions.Same(t1, t2))
}
