package main

import "fmt"

func main() {
   var a int = 100;
 
   /* 判断布尔表达式 */
   if a < 20 {
       fmt.Printf("a 小于 20\n" );
   } else {
       fmt.Printf("a 不小于 20\n" );
   }
   fmt.Printf("a 的值为 : %d\n", a);

}