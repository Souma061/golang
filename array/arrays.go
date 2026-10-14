package main

import "fmt"


func main() {
    // var nums[4] int 
    // nums[1] = 1
    // nums[2] = 2
    nums := [3]int {1,2,3} // to declare in single line
    //2d arrays
    num := [2][2]int {{1,2},{3,4}}
    fmt.Println(nums)
    fmt.Println(num)
}