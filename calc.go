package main

import(

  "fmt"
  "os"
)

func main(){

  var x,y,val float64 = 0.0,0.0,0.0
  var operator string
  fmt.Printf("Enter the first value:  \n")
  fmt.Scanf("%f",&x)
  fmt.Printf("Enter the second value:  \n")
  fmt.Scanf("%f",&y)
  fmt.Print("Enter the operation you would like to perform: \n")
  fmt.Scanf("%s",&operator)
  for true{
    switch operator{
    case "+" :
      val = x+y
      fmt.Printf("%.2f = %.2f + %.2f\n",val,x,y)
    case "-" :
      val = x-y
      fmt.Printf("%.2f = %.2f - %.2f\n",val,x,y)
    case "x" :
      val = x*y
      fmt.Printf("%.2f = %.2f * %.2f\n",val,x,y)
    case "/" :
      val = x/y
      fmt.Printf("%.2f = %.2f / %.2f\n",val,x,y)
    case "q" :
      os.Exit(0)
    default:
      fmt.Printf("Invalid operation.\n")
    }
    fmt.Printf("Enter the operation you would like to perform: \n")
    fmt.Scanf("%s",&operator)
    fmt.Printf("Enter the value: \n")
    fmt.Scanf("%f",&y)
    x = val

  }
}
