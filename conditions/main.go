package main
import "fmt"



func main() {
    age := 11

    if age > 10 && age <= 12 {
        fmt.Println("teenage")
    } else if age > 12 && age <= 18 {
        fmt.Println("bla bla")
    } else {
        fmt.Println("adult")
    }
}