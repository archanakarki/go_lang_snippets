// if I have to return two results I can with swap func
// no need of func name, just swap(x type, y type)
package main

import "fmt"

func swap(x,y string) (newY, newX string){
	newX = y
	newY = x
	return
}

func main(){
	a, b := swap("hello", "world")
	fmt.Println(b, a)
}