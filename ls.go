package main

import  "fmt"
func main(){

    a:=bibao()
    fmt.Println(a())
    fmt.Println(a())
}
//静态数组
func shuzu() {
	var a [10]int
	var i int
	for i=0;i<10;i++ {
		a[i]=i+10
	}
	fmt.Println(a)
}
//切片
func  qiepian() {
	//i:=0
	a := []int{}
	a = append(a, 1, 2, 3)
	a = append(a, 89, 89)
	for k, i := range a {
		println(k, i)
	}
}
	//fmt.Println(a[1:3])
func bibao() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}

	}
