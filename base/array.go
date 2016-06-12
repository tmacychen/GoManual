package main

import "fmt"

func main(){
	fmt.Println("array 测试")

	var a [3] int = [3]int {1,2,3}
	fmt.Println(a[2])
	b := [3] int {1, 2}
	c := [...] int {1,2,3}
	//如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
	fmt.Println("a == b : ", a == b ,"a == c :" , a == c,"b == c :", b == c)


	test := [32]int {1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Printf("test:%p, %v Types :%T\n",&test, test,test)
//	zero(&test)
//	zeroA(test)
	zeroB(&test)
	fmt.Printf("test:%p, %v Types : %T\n",&test, test,test)
}

func zero(ptr *[32] int){
	for i := range ptr{
		ptr[i] = 0
	}
}
func zeroA( ptr [32] int){
	for i := range ptr{
		ptr[i] = 0
	}
	fmt.Printf("in zeroA ** test : %v\n",ptr)
}
func zeroB(ptr * [32]int){
	*ptr = [32] int{}
}