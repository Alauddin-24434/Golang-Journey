package main

import "fmt"

func main() {
    // =============================== Integer ===============================
    fmt.Println("Integer example:", 1+2)

    // =============================== String ================================
    fmt.Println("String example:", "hello go lang")

    // =============================== Boolean ===============================
    fmt.Println("Boolean examples:", true, false)

    // =============================== Float ================================
    fmt.Println("Float example:", 12.4)
}

// =============================== Q&A ===============================

// Q1: Go তে কত ধরনের primitive data type আছে?
// A1: মূলত 4 ধরনের: int, float, string, bool। এছাড়া byte, rune ইত্যাদি আছে।  

// Q2: Integer এবং Float এর মধ্যে পার্থক্য কী?
// A2: Integer হলো সম্পূর্ণ সংখ্যা, Float হলো দশমিক সংখ্যা। যেমন: 12 → int, 12.4 → float64  

// Q3: Boolean type কী এবং কী value নিতে পারে?
// A3: Boolean type শুধুমাত্র দুটি value নিতে পারে: true অথবা false  

// Q4: String কীভাবে declare করা হয়?
// A4: "Double quotes" ব্যবহার করে যেমন: "Hello Go"  

// Q5: Go তে arithmetic operations কোন type এর উপর করা যায়?
// A5: int এবং float type এর উপর arithmetic operation করা যায়।
