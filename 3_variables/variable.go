package main

import "fmt"

func main() {
    // =============================== Variable Declaration ===============================

    // 1. Declare with type explicitly
    var name string = "Alauddin"

    // 2. Declare without type (type inferred automatically)
    var fatherName = "Ab. Aziz"

    // 3. Short declaration (type inferred)
    motherName := "Hamdida"

    fmt.Println("Name:", name)
    fmt.Println("Father Name:", fatherName)
    fmt.Println("Mother Name:", motherName)

    // =============================== Multiple Variables ===============================
    // Declaring multiple variables of the same type
    var a, b, c, d int = 1, 3, 4, 5
    fmt.Println("Multiple variables:", a, b, c, d)

    // Multiple variables with different types
    var myName, age, isAdult, height = "Alauddin", 25, true, 5.6
    fmt.Println(myName, age, isAdult, height)
}

// =============================== Q&A ===============================

// Q1: Go তে variable declare করার বিভিন্ন উপায় কী কী?
// A1: 
//    1. With type explicitly: var name string = "value"
//    2. Type inferred automatically: var name = "value"
//    3. Short declaration: name := "value"

// Q2: এক লাইনে একাধিক variable declare করা যায় কীভাবে?
// A2: var a, b, c int = 1, 2, 3  অথবা ভিন্ন টাইপের জন্য var x, y, z = "A", 25, true

// Q3: Short declaration (:=) কবে ব্যবহার করা যায়?
// A3: যখন ফাংশনের ভিতরে নতুন variable declare করতে হয়। global বা package scope এ ব্যবহার করা যায় না। 

// Q4: Go তে ভেরিয়েবল type auto infer হয় কীভাবে?
// A4: ভ্যালু অনুযায়ী Go compiler type নির্ধারণ করে। যেমন: "Alauddin" → string, 25 → int, 5.6 → float64, true → bool
