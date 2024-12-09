package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	HEADLESS_MODE = true // 控制是否使用无头模式
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		fmt.Println("使用方法: ./程序名 <URL>")
		fmt.Println("示例: ./程序名 https://huihifi.com/evaluation/xxx")
		os.Exit(1)
	}

	url := os.Args[1]

	if strings.Contains(url, "rtings.com") {
		fmt.Println("检测到Rtings链接，使用Rtings处理方法...")
		processRtings(url)
	} else if strings.Contains(url, "huihifi.com") {
		fmt.Println("检测到毁HiFi链接，使用毁HiFi处理方法...")
		processHuiHiFi(url)
	} else {
		fmt.Println("不支持的URL类型")
		os.Exit(1)
	}
}
