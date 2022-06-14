package main

import "fmt"


// for init; cond ; post {

}//

// int v = a+b++-c


func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i++
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}