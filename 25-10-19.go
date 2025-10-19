package main

import "fmt"

func main() {
	var a = 23
	fmt.Println(a)
	b := 24
	fmt.Println(b)
	fmt.Print("用Print：")
	fmt.Print("Hello")
	fmt.Print("world\n")
	fmt.Println("I", "love", "you")
	fmt.Println("hello", "world")
	fmt.Println("tell", "hello")
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", a)
	var name string
	var age int
	fmt.Println(`Input:`)
	_, err := fmt.Scanln(&name, &age)
	if err != nil {
		return
	}
	fmt.Println(name, age)

}
