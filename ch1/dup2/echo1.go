// echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
)

// 一般输入来自一个外部源：文件、网络连接、其他程序的输出、键盘、命令行参数等

func main()  {
	// 声明两个 string 的变量。如果没有明确的初始化，默认为【空值】
	var s, sep string
	// := 用于【短变量声明】，这种语句声明一个或多个变量。并且根据初始化的值给予合适的类型
	for i :=1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}