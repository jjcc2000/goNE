package main

import("fmt")

func main(){
	dataGruping()

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
func ej5(){
	//You can use raw straight variables
	a:= `this is a string that can have 
multiples lines and it also can have this "Look at it" `
	fmt.Println(a)
}
func ej6(){
	const(
		a = 2023 + iota
		b = 2023+iota
		c = 2023+ iota
		d= 2023+iota
	)
	//IOTA still work :)
	fmt.Println("This is the exercise 6")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func dataGruping(){
	//X := type{Values}//Composite Literal
	x := []int {4,5,6,7,8}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Print(x[4]) //Indes Position

	//To loop through Slice
	fmt.Println("To loop through Slice")
	for i, v := range x{
		fmt.Printf("Position: %d, Value: %d\n", i,v )
	}
	//Slicin an Slice 
	fmt.Println("Slice an Slice")
	fmt.Println(x[1:])
	fmt.Println(x[0:2])
	//APend a Slice
	fmt.Println("The y slice is append to the X array")
	y := []int{1212,313,234234,234,45345}
	x = append(x, y...)
	fmt.Println(x)
	//Deleting from a Slice 
	fmt.Println("To Delete from a slice you need to apeend it with the respectives Position")
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}