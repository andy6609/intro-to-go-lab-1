package main

import "fmt"

func main () {
	for i:=0 ; i<10 ; i++ {                // it is not i = 0 . i := 0 is correct 
		fmt.Println("hello, world!") // fmt.print() -> unexported 
		}
}	 