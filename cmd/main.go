/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         main.go
@ Create Time:  2020/4/28 11:09
@ Software:     GoLand
*/

package main

import (
	"fmt"
	. "go-example/models/config"
)

func main() {
	Config.Print()

	fmt.Println("hello world.")
}
