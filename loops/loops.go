package main

import "fmt"

//for -> in go, there is only for loop. we will use it for all the loops like while
func main() {
   // //while style
   // i := 1
   // for i <= 3 {
   //     fmt.Println(i)
   //     i+= 1
   // }
   // infinite loop
   // for {
   //     fmt.Println(i)
   // }
   // classic for loop
   for i := 0;i < 3;i++ {
       // break
       if i == 2 {
           continue
       }
       fmt.Println(i)
   }
    
}
