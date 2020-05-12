// If two or more consecutive function params share type only keep it in the last
package main

import "fmt"

func add(x,y int) int{
	return x + y
}

func main(){
	fmt.Println(add(10, 12))
}