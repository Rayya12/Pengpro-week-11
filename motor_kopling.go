package main

import "fmt"

func main(){
  var gigi int
  var kopling, gas bool
  fmt.Scanln(&gigi,&kopling,&gas)
  if gigi == 0 && kopling && gas{
    fmt.Println("mesin menyala dan motor tidak berjalan")
  } else if gigi != 0 && !kopling && gas{
    fmt.Println("motor berjalan")
  } else if gigi != 0 &&  kopling && !gas{
    fmt.Println("Mesin menyala dan motor tidak berjalan")
  } else{
    fmt.Println("Mesin mati")
  }
}
