package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	hiro := Student {
		name: "Hiro",
		age:  20,
		class: "Computer Science",
		pseudonym: "Geeko",
	}

	hiro.displayDetails()
	min := 12
	max := 45
	randomValue := generateRandomValue(min, max)

	fmt.Println("Random generated value between", min, "and", max, "is", randomValue)
	fmt.Printf("The square root of the generated value is %.3v", calculateSquareRoot(randomValue))

	fmt.Println("The factoriel of 7 is ", factoriel(7))

	haojue := createProduct("Haojue", "A spidy motorcycle", 700000.0)
	purse := getPurse("Gogo Purse", "Black", 1000000.0)

	purse.buyProduct(haojue)
	writeToConsole("Hi, i'am using fprintf")
	showOS()
	fmt.Println("Random marks", getRandomBinary(6))
}

func calculateSquareRoot(randomValue int) float64 {
	return math.Sqrt(float64(randomValue))
}


/// Create a student type
type Student struct {
	name string
	age int
	class string
	pseudonym string
}

func (s *Student) displayDetails() {
	fmt.Println("\nHere are the details of the student with pseudo:", s.pseudonym)
	fmt.Println("Name: ", s.name)
	fmt.Println("Age: ", s.age)
	fmt.Println("Level: ", s.class)
}

func generateRandomValue(min, max int) int {
	return min + rand.Intn(max-min)
}

func calculateMean(numbers []int) float64 {
	sum := 0
	total := len(numbers)
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(total)
}

/// Recursively calculate the factoriel of a number
func factoriel(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factoriel(n-1)
}

/// Recursively calculate the fibonacci sequence of a number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/// Create a purse type
type Purse struct {
	name string
	color string
	amount float64
}

func getPurse(name, color string, amount float64) Purse {
	return Purse{
		name: name,
		color: color,
		amount: amount,
	}
}

/// Create a product type
type Product struct {
	name string
	description string
	price float64
}

func createProduct(name, description string, price float64) Product {
	return Product{
		name: name,
		description: description,
		price: price,
	}
}

func (p *Purse) buyProduct(product Product) {
	if p.amount < product.price {
		fmt.Println("You don't have enough money to buy this product.")
	} else {
		p.amount -= product.price
		fmt.Println("You have bought", product.name, "for", product.price, "dollars.")
		fmt.Println("Your purse amount is now", p.amount, "dollars")
	}
}

func writeToConsole(content string) {
	lines, err := fmt.Fprintln(os.Stdout, content)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing to console:", err)
	}
	fmt.Println("Wrote", lines, "bytes to console")
}

func showOS() {
	switch osys := runtime.GOOS; osys {
	case "linux":
		fmt.Println("Linux OS")
	case "darwin":
		fmt.Println("Darwin OS")
	case "windows":
		fmt.Println("Windows OS")
	case "freebsd":
		fmt.Printf("FreeBSD OS")
	}
}

func greeting() {
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func getRandomBinary(number int) []int {
	bins := make([]int, 0)
	for i := 0; i < number; i++ {
		bins = append(bins, rand.Intn(2))
	}
	return bins
}