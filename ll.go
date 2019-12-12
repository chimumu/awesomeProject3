package main

import (
	"fmt"
	"net/http"
)



func  main() {
	/*
	_,f,k:=mas()
	fmt.Println(f,k)

	var sum int
	sum=mercury()
	fmt.Printf("面积为：%d",sum)

	 */
	httpd()
}

//比较三个数的大小按照由小到大输出
func  bijiaodaxiao() {
	var  x,y,z,c float64
	fmt.Println("请输入三个数:")
	fmt.Scan(&x,&y,&z)
    if x<y {
    	c=x
    	x=y
    	y=c
	}
	if x<z {
		c=z
		z=x
		x=c
	}
	if y<z{
		c=z
		z=y
		y=c
	}

	fmt.Println(z,y,x)

}

//变量
func mas() (int, int, string){
	a, b, c :=1, 2, "str"
	return a,b,c
}

//常量
func mercury() (int){
	const chang=100
	const kuan=2
	var area int
	const a,b,c=1,false,"str"
	area=chang*kuan
	return area
}

//计算
func config() {
	a:=10
	b:=20
	c:=0
	if (a<b) {
		c = a + b
		fmt.Printf("第一行为加：%d\n", c)
	}else{
		c = a - b
		fmt.Printf("第二行为减：%d\n", c)
	}

}


//net/http库的使用
func httpd() {
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080",nil)
}
