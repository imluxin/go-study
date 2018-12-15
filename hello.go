package main

// 导入包
import (
	"fmt"
	// "os"
	// "io"
)

// 定义常量
const PI = 3.14

// 全局变理的声明与赋值
var name = "gopher"
var (
	v1 = "123"
	v2 = "456"
	v3 = "789"
	v4 = "000"
)

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由 main 函数作为程序入口点启动
func main() {
	fmt.Println("Hello go! \n" + name)
	fmt.Println(v1 + v2 + v3)
}