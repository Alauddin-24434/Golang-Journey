package main

import "fmt"


func main(){


//  with type string declare
	var name string = "alauddin"
	// without type but infer type string auto
	var fathername ="Ab. aziz"
    // type is  infared
	motherName := "hamdida"

	fmt.Println("name:", name)
	fmt.Println("fathername:", fathername)
	fmt.Println("motherName:", motherName)

	// mutiple variable declare 
	var a, b, c, d int= 1, 3, 4,5

		fmt.Println("muliple variable:", a,b,c,d)

		var myName , age , isAdult, hight = "Alauddin", 25, true, 5.6;

		fmt.Println(myName,age, isAdult, hight)

}