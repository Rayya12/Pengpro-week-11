package main

import "fmt"

func main(){
  var minggu int
  var n, uang, total int

  total = 0
  fmt.Scanln(&minggu)
  if minggu == 4{
    n = 4
  } else {
    n = 5
  }
  for i := 1 ; i <= n; i++{
    fmt.Scanln(&uang)
    total += uang
  }
  fmt.Println(total)
}
