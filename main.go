package main

import (
	"fmt"
	// "math"
	// "math/rand"
	// "rsc.io/quote"
)
func add(x int ,y int) (int,int){
  return x,y
}
func mul1(x,y float64) float64{
  return x*y;
}
func mul(x,y float32) float32{
  return x*y;
}
func main() {
    // fmt.Println(quote.Glass())
    // fmt.Println(quote.Go())
    // fmt.Println(quote.Hello())
    // fmt.Println(quote.Opt())
    // fmt.Println(rand.Intn(10))
    // fmt.Println(math.Ceil(3.8906))
    // fmt.Println(math.Ceil(math.Pi))
  fmt.Println(add(2,3))
  fmt.Println(mul(3.583,4.095674))
  fmt.Println(mul1(3.583,4.095674))
  var  a=10;
  fmt.Printf("%v",a);
  
}