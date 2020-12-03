package main

import "fmt"

func main(){
	fmt.Printf("%d", gcd(25, 20) )
}
// WOW
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
