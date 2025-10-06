package main

import "fmt"

// =============================== Constants ===============================
// Declaring a constant
const myFathername = "Abdul Aziz"

func main() {
    // Printing the constant
    fmt.Println("My father's name is:", myFathername)

    // Printing an empty line
    fmt.Println()
}

// =============================== Q&A ===============================

// Q1: Go তে constant কি?
// A1: Constant হলো এমন ভেরিয়েবল যার value পরিবর্তন করা যায় না। এটি compile time এ নির্ধারিত হয়।  

// Q2: constant declare করার syntax কীভাবে?
// A2: const name = value  (optional type: const name type = value)  

// Q3: Go তে কোন কোন type constant হতে পারে?
// A3: Numeric, string, boolean constants হতে পারে। যেমন:
//     const pi = 3.1415
//     const greeting = "Hello"
//     const isTrue = true

// Q4: Variable vs Constant পার্থক্য কী?
// A4: 
//    - Variable পরিবর্তনযোগ্য, Constant অপরিবর্তনীয়।  
//    - Variable declare করতে var ব্যবহার হয়, Constant declare করতে const ব্যবহার হয়।  
