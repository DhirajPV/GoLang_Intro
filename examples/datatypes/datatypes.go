package main

import "fmt"

func main()  {
	fmt.Println("go" + "lang");

	fmt.Println("1+1=", 1+1);
	fmt.Println("5.0/3.0=", 5.0/3.0);
	// fmt.Println("divideByZero = ", 5.0/0.0);

	fmt.Println(true && false);
	fmt.Println(true || false);
	fmt.Println(!true);

	var a = "initial"
	fmt.Println(a)
	
	var b, c int = 1, 2
	fmt.Println(b, c)
	
	var d = true
	fmt.Println(d)
	
	var e int
	fmt.Println(e)
	
	f := "apple"
    fmt.Println(f)
}