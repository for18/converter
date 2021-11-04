/*
 * @Author: your name
 * @Date: 2021-11-04 14:57:45
 * @LastEditTime: 2021-11-04 16:48:59
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /converter/main.go
 */
//Author:TruthHun
//Email:TruthHun@QQ.COM
//Date:2018-01-21
package main

import (
	"os"

	"fmt"

	"github.com/jmt99/converter/converter"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("错误：缺少指定的json配置文件")
	} else {
		if converter, err := converter.NewConverter(args[0]); err != nil {
			fmt.Println(err.Error())
		} else {
			if err = converter.Convert(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
