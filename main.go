package main

import ("fmt" 
		"github.com/learning-go-book-2e/ch02")


func main(){
	x := complex(2.5,3.1)
	y := complex(10.2,2)
	fmt.Println(x+y)
	fmt.Println(x-y)
	fmt.Println(x*y)
	fmt.Println(x/y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

}

 	//Chapter #1
// {
// 	var flag bool 
// 	var isAwesome = true

// 	var a,b uint16
// }


// func main() {
// 	fmt.Println("Hello, world!")
// }