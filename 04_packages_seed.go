//The program:
// 1) starts running package main
// 2) by using import paths "fmt" and "math/rand" the program is using mentioned pkgs
// 3) in import, pkg name is last element of import path like
// --- "math/rand" means math package includes file that begins with line/statement Package rand


//rand.Seed used
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite num is", rand.Intn(10))
}