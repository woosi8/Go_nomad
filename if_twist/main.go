package main

import "fmt"

func canIDrink(age int) bool{
	if koreanAge := age +2 ; koreanAge < 18 { //if에 바로 변수를 선언해서 쓰는이유는 이렇게 함으로써 if에 쓰기위한 변수라는 목적을 내포하기위함
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(14))
}


