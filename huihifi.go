package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/chromedp/chromedp"
)

// 定义数据结构
type FreqPoint struct {
	Freq float64 `json:"freq"`
	SPL  float64 `json:"spl"`
}

func processHuiHiFi(url string) {
	fmt.Printf("目标URL: %s\n", url)
	fmt.Printf("无头模式: %v\n", HEADLESS_MODE)
	fmt.Println("--------------------")
	fmt.Printf("开始采集数据...\n网页正在加载中，请耐心等待...\n")

	opts := chromedp.DefaultExecAllocatorOptions[:]
	if !HEADLESS_MODE {
		opts = append(opts, chromedp.Flag("headless", false))
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var lastFreq, lastSpl string
	freqRegex := regexp.MustCompile(`Freq\(Hz\)：([\d.]+)`)
	splRegex := regexp.MustCompile(`SPL\(dB\)：([\d.]+)`)

	var dataPoints []FreqPoint
	var productName string

	err := chromedp.Run(ctx,
		chromedp.EmulateViewport(2160, 1440),
		chromedp.Navigate(url),
		chromedp.WaitVisible("canvas"),
		chromedp.Text("h2", &productName),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var canvasBox struct {
				X      float64 `js:"x"`
				Y      float64 `js:"y"`
				Width  float64 `js:"width"`
				Height float64 `js:"height"`
			}

			err := chromedp.Evaluate(`
				(() => {
					const canvas = document.querySelector('canvas');
					const rect = canvas.getBoundingClientRect();
					return {
						x: rect.x,
						y: rect.y,
						width: rect.width,
						height: rect.height
					};
				})()
			`, &canvasBox).Do(ctx)

			if err != nil {
				return err
			}

			fmt.Printf("Canvas dimensions: x=%f, y=%f, width=%f, height=%f\n",
				canvasBox.X, canvasBox.Y, canvasBox.Width, canvasBox.Height)

			totalPoints := int(canvasBox.Width)
			currentPoint := 0

			for x := float64(0); x < canvasBox.Width; x++ {
				currentPoint++
				if currentPoint%10 == 0 {
					fmt.Printf("\033[2K\r")
					fmt.Printf("进度: [")
					progress := int((float64(currentPoint) / float64(totalPoints)) * 50)
					for i := 0; i < 50; i++ {
						if i < progress {
							fmt.Print("#")
						} else {
							fmt.Print(" ")
						}
					}
					fmt.Printf("] %.1f%% (%d/%d)\n",
						float64(currentPoint)/float64(totalPoints)*100,
						currentPoint,
						totalPoints)
				}

				err := chromedp.MouseClickXY(
					canvasBox.X+x,
					canvasBox.Y+canvasBox.Height*0.2,
				).Do(ctx)

				if err != nil {
					continue
				}

				var allText string
				err = chromedp.Evaluate(`document.body.innerText`, &allText).Do(ctx)
				if err != nil {
					continue
				}

				freqMatches := freqRegex.FindStringSubmatch(allText)
				splMatches := splRegex.FindStringSubmatch(allText)

				if len(freqMatches) > 1 && len(splMatches) > 1 {
					freq := freqMatches[1]
					spl := splMatches[1]

					if freq != lastFreq || spl != lastSpl {
						fmt.Printf("\033[2K\r")
						fmt.Printf("\rFreq(Hz)：%s, SPL(dB)：%s", freq, spl)

						lastFreq = freq
						lastSpl = spl

						if freqVal, err := strconv.ParseFloat(freq, 64); err == nil {
							if splVal, err := strconv.ParseFloat(spl, 64); err == nil {
								dataPoints = append(dataPoints, FreqPoint{
									Freq: freqVal,
									SPL:  splVal,
								})
							}
						}
					}
				}
			}

			fmt.Printf("\033[2K\r")
			fmt.Printf("进度: [##################################################] 100%% (%d/%d)\n",
				totalPoints, totalPoints)
			fmt.Println("\n扫描完成！")
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	productName = regexp.MustCompile(`[\\/:*?"<>|]`).ReplaceAllString(productName, "_")

	// 直接保存为CSV文件
	csvFilename := fmt.Sprintf("%s.csv", productName)
	csvFile, err := os.Create(csvFilename)
	if err != nil {
		log.Fatal("创建CSV文件失败:", err)
	}
	defer csvFile.Close()

	// 添加 UTF-8 BOM
	csvFile.Write([]byte{0xEF, 0xBB, 0xBF})

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, point := range dataPoints {
		writer.Write([]string{
			strconv.FormatFloat(point.Freq, 'f', 2, 64),
			strconv.FormatFloat(point.SPL, 'f', 2, 64),
		})
	}

	fmt.Printf("数据已保存到CSV文件: %s\n", csvFilename)
	fmt.Printf("总计采集数据点: %d\n", len(dataPoints))
}
