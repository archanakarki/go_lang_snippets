// if I have to return two results I can with swap func
// no need of func name, just swap(x type, y type)
package main

import "fmt"

func swap(x,y string) (string, string){
	return y, x
}

func main(){
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}