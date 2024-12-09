package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/chromedp/chromedp"
)

func processRtings(url string) {
	fmt.Printf("正在启动浏览器...\n")
	fmt.Printf("目标URL: %s\n", url)
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
		chromedp.Evaluate(`
			(() => {
				const tables = document.getElementsByTagName('table');
				if (tables.length > 0) {
					const rows = tables[0].getElementsByTagName('tr');
					let csvContent = '';
					for (let i = 2; i < rows.length; i++) {
						const cells = rows[i].getElementsByTagName('td');
						// 检查是否有至少3列，且第3列不为空
						if (cells.length >= 3 && cells[2].textContent.trim() !== '') {
							// 只取第1列和第3列的值
							const col1 = cells[0].textContent.replace(/,/g, '').trim();
							const col3 = cells[2].textContent.replace(/,/g, '').trim();
							csvContent += col1 + ',' + col3 + '\n';
						}
					}
					return csvContent;
				}
				return '';
			})()
		`, &tableContent),
	)

	if err != nil {
		log.Fatal(err)
	}

	// 处理文件名（移除非法字符）
	chartTitle = regexp.MustCompile(`[\\/:*?"<>|]`).ReplaceAllString(chartTitle, "_")
	csvFilename := chartTitle + ".csv"

	// 创建CSV文件
	csvFile, err := os.Create(csvFilename)
	if err != nil {
		log.Fatal("创建CSV文件失败:", err)
	}
	defer csvFile.Close()

	// 直接写入CSV内容
	_, err = csvFile.WriteString(tableContent)
	if err != nil {
		log.Fatal("写入CSV失败:", err)
	}

	fmt.Printf("数据已保存到: %s\n", csvFilename)
}
