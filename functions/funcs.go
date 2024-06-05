package functions

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)


func WorldCount(s string) map[string]int {
  res := make(map[string]int)
  for _, word := range strings.Fields(s) {
    res[word]++
  }
  return res
}

func RandomBinary(number int) []int {
  bins := make([]int, 0)
  for i := 0; i < number; i++ {
    bins = append(bins, rand.Intn(2))
  }
  return bins
}


func CalculateSquareRoot(randomValue int) float64 {
  return math.Sqrt(float64(randomValue))
}

func GenerateRandomValue(min, max int) int {
  return min + rand.Intn(max-min)
}

func CalculateMean(numbers []int) float64 {
  sum := 0
  total := len(numbers)
  for _, num := range numbers {
    sum += num
  }
  return float64(sum) / float64(total)
}

// / Recursively calculate the factoriel of a number
func Factoriel(n int) int {
  if n == 0 || n == 1 {
    return 1
  }
  return n * Factoriel(n-1)
}

// / Recursively calculate the fibonacci sequence of a number
func Fibonacci(n int) int {
  if n <= 1 {
    return n
  }
  return Fibonacci(n-1) + Fibonacci(n-2)
}



func WriteToConsole(content string) {
  lines, err := fmt.Fprintln(os.Stdout, content)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error writing to console:", err)
  }
  fmt.Println("Wrote", lines, "bytes to console")
}


/// Determine the OS type of the current machine
func OS() {
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

func RigthGreeting() {
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