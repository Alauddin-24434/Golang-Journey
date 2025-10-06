package main

import "fmt"

func main() {
    // =============================== Array Examples ==============================
    // Arrays with defined lengths
    var numbers = [3]int{1, 2, 3}
    fmt.Println(numbers)

    country := [5]string{"Bangladesh", "Nepal", "Bhutan", "India", "Pakistan"}
    fmt.Println(country)

    // Arrays with inferred lengths
    var age = [...]int{25, 36, 32, 78, 67, 8}
    fmt.Println(age)

    roll := [...]int{2, 4, 6, 7, 8}
    fmt.Println(roll)

    // =============================== Access Elements ==============================
    prices := [3]int{10, 20, 30}
    fmt.Println(prices[0]) // First element
    fmt.Println(prices[2]) // Third element

    // =============================== Change Elements ==============================
    prices[2] = 50
    fmt.Println(prices) // Updated array

    // =============================== Array Initialization =========================
    arr1 := [5]int{}             // Not initialized, default 0
    arr2 := [5]int{1, 2, 3}      // Partially initialized
    arr3 := [5]int{1, 2, 3, 4, 5} // Fully initialized
    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(arr3)

    // =============================== Only Specific Elements =======================
    arr4 := [5]int{0: 20, 3: 50} // Initialize only 0th and 3rd index
    fmt.Println(arr4)

    // =============================== Find Length of Array =========================
    arr5 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
    arr6 := [...]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println(len(arr5))
    fmt.Println(len(arr6))
}
