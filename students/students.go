package students

import "fmt"

/// Create a student type
type Student struct {
  Name      string
  Age       int
  Class     string
  Pseudonym string
}

func Create(pseudonym, name, class string, age int) Student {
  return Student{Name: name, Age: age, Class: class, Pseudonym: pseudonym}
}

func (s Student) DisplayDetails() {
  fmt.Println("\nHere are the details of the student with pseudo:", s.Pseudonym)
  fmt.Println("Name: ", s.Name)
  fmt.Println("Age: ", s.Age)
  fmt.Println("Level: ", s.Class)
}