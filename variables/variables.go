package main

import "fmt"


func main() {
    // var str1 string = "golang"
    var str1 = "golang" // automatically infers the type if we do not declare the type
    //var isAdult bool = true
    var isAdult = false // same as above

    //shorthand syntax
    name := "souma"
    // var name string
    // name = "souma"
    fmt.Println(str1)
    fmt.Println(isAdult)
    fmt.Println(name)
}