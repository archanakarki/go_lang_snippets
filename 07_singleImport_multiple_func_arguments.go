// Single import statement
// Two arguments in function, each has name and int type comma other arguments

package main

import "fmt"

func add(x int, y int) int{
	return x + y
}

func main(){
	fmt.Println(add(10, 12))
}