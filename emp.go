package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp6 := Employee{
        firstName: "Divya",
        lastName:  "Avula",
        age:       21,
        salary:    25000,
    }
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d\n", emp6.salary)
    emp6.salary = 6500
    fmt.Printf("New Salary: $%d", emp6.salary)
}
