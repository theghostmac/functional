package main

import(
    "fmt"
)

func countdown(number int) {
    if number == 0 {
       fmt.Println("Blast off!")
       return
    }
    fmt.Println(number)
    countdown(number - 1)
}

func main() {
    countdown(10)
}
