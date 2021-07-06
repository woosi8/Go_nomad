package main

import (
	"fmt"
	"strings"
)


func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done") //함수 리턴이 끝나고 추가하고 싶을때
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return

}


func main() {
	totalLenght, _ := lenAndUpper("nico")
	fmt.Println(totalLenght)
	totalLenght, up := lenAndUpper("nico")
	fmt.Println(totalLenght, up)
}