package main

import (
    "fmt"
)

func Swap(a []int, i int){
    if (i >= 0) && ((i + 1) <  cap(a)) {
        var temp int
        temp = a[i]
        a[i] = a[i + 1]
        a[i + 1] = temp
    }
}

func BubbleSort(a []int){
    for i := len(a) - 1; i >= 0; i-- {
        for j := 0; j < i; j++ {
            if a[j] > a[j + 1] {
                Swap(a, j)
            }
        }
    }
}

func main() {
    var a [10]int

    fmt.Println("Enter 10 integers separated by space: ")
    for i := 0; i < 10; i++  {
        fmt.Scanf("%d", &a[i])
    }
    fmt.Println("User input:")
    fmt.Println(a)
    BubbleSort(a[:])
    fmt.Println("Sorted (ascending) output:")
    fmt.Println(a)
}