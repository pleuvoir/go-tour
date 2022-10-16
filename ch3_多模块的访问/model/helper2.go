package model

import "fmt"

func Say2() {
	fmt.Printf("%s\n", "准备输入..")
	fmt.Printf("版本号=%s", appVersion)
	fmt.Printf("%s\n", "结束输入..")
}
