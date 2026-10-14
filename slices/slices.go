package main

import "fmt"

//slices -> dynamic
func main() {
    var nums = make([]int,2,5)
    fmt.Println(cap(nums))
    for i := 1; i <= 6; i++ {
        nums = append(nums, i)
        fmt.Println(len(nums), cap(nums))
    }

    a := make([]int, 2, 5)
    b := a
    a = append(a, 1)   // fits, same array
    b = a[:2]
    a[0] = 99
    fmt.Println(b[0])  // 99, shared
    
    a = append(a, 2, 3, 4) // overflows, new array
    a[0] = 7
    fmt.Println(b[0])  // still 99, b points to the old array
    fmt.Println(a[0])

}