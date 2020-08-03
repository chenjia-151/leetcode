package main  // 声明 包   模块化 module
//  go 是一个天生支持多核计算的云开发时代C语言
// node + go 两门语言  服务器有关
// require('fs')  fileSystem
import(   //es6
	"fmt"  // 向命令行输出hello world   format
)
// node 脚本语言
// go  和  c  一样   二进制语言
// 编译器  func => function
func minEatingSpeed(piles []int, H int) int {
	low := 1  // let low = 1;   声明一个变量
	// 最大pile的数量  11  交给一个函数
	// 函数是组织代码的最小单元 { 块级作用域 }
	hi := maxPiles(piles)
	// fmt.Printf("%d",hi)
	for{
		// go 没有while 
		if low > hi{
			break;
		}
		if canEatAllBanas(piles,H,low){
			return low
		}else{
			low++
		}
	}
	// // 是否可以吃完  1..max low
	return low
}

func canEatAllBanas(piles []int, H int, k int)bool{
	// 算法
	// 1. range piles piles[i]/k 向上取整数  Math.cell  对浮点数向上取整
	// 2. 加起来
	// 3. 是否小于等于 H
	return true
}

func maxPiles(piles []int) int {
	h := 0
	// range 
	for _, n:= range piles{
		h = max(h,n)
	}
	return h
}
func max(a int ,b int ) int {
	if a > b {
		return a
	}
	return b
	// return  a > b ? a : b
}

func main(){  //  func 函数关键字  入口函数
	// 从 main 开始
	// go 是一个静态的  类型约束
	fmt.Printf("%d小时之内吃完香蕉的最慢速度是%d只/小时",8,minEatingSpeed([]int{3,6,7,11},8))  //传一个数组，并声明数组的类型
}

// go run main.go
// pwd
// dir 