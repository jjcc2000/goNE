package main

import("fmt")

func main(){
	ej1()
	ej2()
	ej4()
}

func ej1(){
	fmt.Println("Exercise 1")
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
}
func ej2(){
	fmt.Println("Exercise 2")
	a:= (10>9)
	b:= (10==10)
	c:= (10>11)
	d:= (10!=10)

	fmt.Println(a,b,c,d)
}
func ej4(){
	fmt.Println("Exercise 4")
	a:= 42
	fmt.Printf("%d\t%b\t%#x\n",a,a,a,)

	fmt.Println("After bit swiching")
	b:= a<<1
	fmt.Printf("%d\t%b\t%#x\n",b,b,b)
}