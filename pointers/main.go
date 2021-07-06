package main

import "fmt"



func main() {
	a := 2
	b := &a //pointer하면 같은 주소를 바라보기 때문에 복사가 아니라 영향을 받아 같아진다
	a = 10
	*b = 20	//b를 이용해 a값 변경하기 , 위에서 b가 a의 주소를 공유하고(pointer) 있기때문에 가능
	fmt.Println(a,*b,&a)
}