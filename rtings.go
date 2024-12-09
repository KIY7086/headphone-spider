package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func processRtings(url string) {
	// 处理URL
	baseURL := strings.Split(url, "?")[0]      // 删除?及后面的内容
	url = baseURL + "?disabled=0:1:,0:2:,0:3:" // 添加新的参数

	fmt.Printf("正在启动浏览器...\n")
	fmt.Printf("处理后的URL: %s\n", url)
	fmt.Println("--------------------")

	opts := chromedp.DefaultExecAllocatorOptions[:]
	if !HEADLESS_MODE {
		opts = append(opts, chromedp.Flag("headless", false))
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var tableContent string
	var chartTitle string
	err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(url),
		chromedp.WaitVisible("rect", chromedp.ByQuery),
		// 获取图表标题
		chromedp.Text(".graph_tool_legend_section-header", &chartTitle, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Println("正在定位rect元素...")
			// 获取rect元素的位置
			var rectBox struct {
				X      float64 `js:"x"`
				Y      float64 `js:"y"`
				Width  float64 `js:"width"`
				Height float64 `js:"height"`
			}

			err := chromedp.Evaluate(`
				(() => {
					const rect = document.querySelector('rect');
					const bbox = rect.getBoundingClientRect();
					return {
						x: bbox.x,
						y: bbox.y,
						width: bbox.width,
						height: bbox.height
					};
				})()
			`, &rectBox).Do(ctx)

			if err != nil {
				return err
			}

			// 移动鼠标到rect中心
			return chromedp.MouseClickXY(
				rectBox.X+rectBox.Width/2,
				rectBox.Y+rectBox.Height/2,
			).Do(ctx)
		}),
		// 提取表格内容
		chromedp.Evaluate(`
			(() => {
				const tables = document.getElementsByTagName('table');
				if (tables.length > 0) {
					const rows = tables[0].getElementsByTagName('tr');
					let result = [];
					for (let i = 2; i < rows.length; i++) {
						const cells = rows[i].getElementsByTagName('td');
						if (cells.length >= 2) {
							let freq = cells[0].textContent.replace(/,/g, '');
							let value = cells[1].textContent.replace(/,/g, '');
							result.push([freq, value]);
						}
					}
					return JSON.stringify(result);
				}
				return '[]';
			})()
		`, &tableContent),
	)

	if err != nil {
		log.Fatal(err)
	}

	// 处理文件名（移除非法字符）
	chartTitle = regexp.MustCompile(`[\\/:*?"<>|]`).ReplaceAllString(chartTitle, "_")
	csvFilename := chartTitle + ".csv"

	// 解析JSON数据
	var data [][]string
	err = json.Unmarshal([]byte(tableContent), &data)
	if err != nil {
		log.Fatal("解析表格数据失败:", err)
	}

	// 创建CSV文件
	csvFile, err := os.Create(csvFilename)
	if err != nil {
		log.Fatal("创建CSV文件失败:", err)
	}
	defer csvFile.Close()

	// 创建CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// 写入数据
	for _, row := range data {
		if err := writer.Write(row); err != nil {
			log.Fatal("写入CSV失败:", err)
		}
	}

	fmt.Printf("数据已保存到: %s\n", csvFilename)
}
