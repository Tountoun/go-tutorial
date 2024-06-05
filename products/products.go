package products


/// Create a product type
type Product struct {
  Name        string
  Description string
  Price       float64
}


func New(name, description string, price float64) Product {
  return Product {
    Name:        name,
    Description: description,
    Price:       price,
  }
}
