package clients

import (
  "fmt"
  "main/products"
)

/// Create a client type
type Client struct {
	Name string
	Age int
	Address string
  Purse *Purse
}

/// Create a purse type
type Purse struct {
  Name   string
  Color  string
  Amount float64
}

func NewPurse(name, color string, amount float64) Purse {
  return Purse{Name: name, Color: color, Amount: amount}
}

func New(name, address string, age int) Client {
  return Client{Name: name, Address: address, Age: age,}
}


func (client *Client) SetPurseUsingProperties(name, color string, amount float64) {
  purse := Purse{Name: name, Color: color, Amount: amount}
  client.Purse = &purse
}

func (client *Client) SetPurse(purse *Purse) {
  client.Purse = purse
}


func (client *Client) BuyProduct(product products.Product) {
  if client.Purse != nil {
    purseAmount := client.Purse.Amount
    if purseAmount < product.Price {
      fmt.Println("You don't have enough money to buy this product.")
    } else {
      client.Purse.Amount -= product.Price
      fmt.Println("Client", client.Name, "bought", product.Name, "for", product.Price, "dollars.")
      fmt.Println("His purse amount is now", client.Purse.Amount, "dollars")
    }
  } else {
    fmt.Printf("Client %s doesn't have a purse.\n", client.Name)
  }
  
}