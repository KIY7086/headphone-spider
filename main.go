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
		fmt.Println("使用方法: ./headphone-spider <URL>")
		fmt.Println("支持的URL类型:")
		fmt.Println("1. https://huihifi.com/evaluation/xxx")
		fmt.Println("2. https://rtings.com/headphones/reviews/xxx")
		fmt.Println("3. https://squig.link/xxx 或其他类似网站")
		fmt.Println("示例: ./headphone-spider https://huihifi.com/evaluation/5e14542b-be71-49e8-add2-d6177bf900dc")
		os.Exit(1)
	}

	url := os.Args[1]

	if strings.Contains(url, "rtings.com") {
		fmt.Println("检测到Rtings链接，使用Rtings处理方法...")
		processRtings(url)
	} else if strings.Contains(url, "huihifi.com") {
		fmt.Println("检测到毁HiFi链接，使用毁HiFi处理方法...")
		processHuiHiFi(url)
	} else if strings.Contains(url, "share=") {
		// 只要URL中包含share参数，就使用squiglink处理方法
		fmt.Println("检测到带有share参数的链接，使用Squig.Link处理方法...")
		processSquigLink(url)
	} else {
		fmt.Println("不支持的URL类型")
		os.Exit(1)
	}
}
