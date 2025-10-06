package main

import "fmt"

func main() {
    // =============================== Array Examples ==============================
    // 1. Arrays with defined lengths
    var numbers = [3]int{1, 2, 3}
    fmt.Println("Numbers array:", numbers)

    country := [5]string{"Bangladesh", "Nepal", "Bhutan", "India", "Pakistan"}
    fmt.Println("Country array:", country)

    // 2. Arrays with inferred lengths
    var age = [...]int{25, 36, 32, 78, 67, 8}
    fmt.Println("Age array:", age)

    roll := [...]int{2, 4, 6, 7, 8}
    fmt.Println("Roll array:", roll)

    // =============================== Access Elements ==============================
    prices := [3]int{10, 20, 30}
    fmt.Println("First price:", prices[0])
    fmt.Println("Third price:", prices[2])

    // =============================== Change Elements ==============================
    prices[2] = 50
    fmt.Println("Updated prices array:", prices)

    // =============================== Array Initialization =========================
    arr1 := [5]int{}               // Not initialized, default 0
    arr2 := [5]int{1, 2, 3}        // Partially initialized
    arr3 := [5]int{1, 2, 3, 4, 5}  // Fully initialized
    fmt.Println("arr1:", arr1)
    fmt.Println("arr2:", arr2)
    fmt.Println("arr3:", arr3)

    // =============================== Only Specific Elements =======================
    arr4 := [5]int{0: 20, 3: 50}   // Initialize only 0th and 3rd index
    fmt.Println("arr4:", arr4)

    // =============================== Find Length of Array =========================
    arr5 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
    arr6 := [...]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println("Length of arr5:", len(arr5))
    fmt.Println("Length of arr6:", len(arr6))
}

// =============================== Q&A ===============================

// Q1: Go তে array কী?
// A1: Array হলো fixed-size ordered collection একই type এর elements এর।  

// Q2: Array declare করার বিভিন্ন উপায় কী কী?
// A2: 
//    1. Fixed size: var arr [5]int
//    2. Inferred size: arr := [...]int{1,2,3}
//    3. Short declaration: arr := [3]string{"a","b","c"}

// Q3: Array element access ও modify কিভাবে করা হয়?
// A3: 
//    Access: arr[0]  
//    Modify: arr[2] = 50

// Q4: Array length কিভাবে বের করা যায়?
// A4: len(arr) ফাংশন ব্যবহার করে।  

// Q5: Partial বা specific element initialization কিভাবে করা যায়?
// A5: arr := [5]int{0:20, 3:50} → শুধু 0th এবং 3rd index initialize করা হয়েছে।  

// Q6: Array এবং Slice এর মধ্যে পার্থক্য কী?
// A6: 
//    - Array fixed-size, slice dynamic-size।  
//    - Slice সহজে grow এবং shrink করা যায়, array নয়।
