package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func processSquigLink(urlStr string) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal("无法解析URL:", err)
	}

	// 获取 share 参数
	share := parsedURL.Query().Get("share")
	if share == "" {
		log.Fatal("URL中缺少share参数")
	}

	var selectedModel string
	// 如果share参数包含逗号，让用户选择
	if strings.Contains(share, ",") {
		// 过滤掉 Custom_Tilt
		var models []string
		for _, model := range strings.Split(share, ",") {
			if !strings.Contains(model, "Custom_Tilt") {
				models = append(models, model)
			}
		}

		if len(models) == 0 {
			log.Fatal("没有找到有效的耳机型号")
		} else if len(models) == 1 {
			selectedModel = models[0]
		} else {
			fmt.Println("请选择耳机型号：")
			for i, model := range models {
				fmt.Printf("%d. %s\n", i+1, strings.ReplaceAll(model, "_", " "))
			}

			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("请输入序号 (1-", len(models), "): ")
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("输入错误，请重试")
					continue
				}

				input = strings.TrimSpace(input)
				choice := 0
				_, err = fmt.Sscanf(input, "%d", &choice)
				if err != nil || choice < 1 || choice > len(models) {
					fmt.Println("无效的选择，请重试")
					continue
				}

				selectedModel = models[choice-1]
				break
			}
		}
	} else {
		if strings.Contains(share, "Custom_Tilt") {
			log.Fatal("不支持 Custom_Tilt 选项")
		}
		selectedModel = share
	}

	// 提取主机名和路径
	host := parsedURL.Host
	urlPath := parsedURL.Path
	if urlPath == "" {
		urlPath = "/"
	}

	// 构建数据文件URL，保持原始路径
	baseURL := fmt.Sprintf("https://%s%s", host, path.Dir(urlPath))
	dataURL := fmt.Sprintf("%s/data/%s L.txt", baseURL, strings.ReplaceAll(selectedModel, "_", " "))

	fmt.Printf("正在下载数据文件: %s\n", dataURL)

	// 下载文件
	resp, err := http.Get(dataURL)
	if err != nil {
		log.Fatal("下载文件失败:", err)
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode == 404 {
		log.Fatal("找不到数据文件，无法解析该网站")
	} else if resp.StatusCode != 200 {
		log.Fatal("服务器返回错误:", resp.Status)
	}

	// 检查第一行数据格式
	scanner := bufio.NewScanner(resp.Body)
	if !scanner.Scan() {
		log.Fatal("文件为空")
	}
	firstLine := scanner.Text()
	isCSV := strings.Contains(firstLine, ",")

	// 重新下载文件（因为已经读取了第一行）
	resp, err = http.Get(dataURL)
	if err != nil {
		log.Fatal("下载文件失败:", err)
	}
	defer resp.Body.Close()

	// 再次检查HTTP状态码
	if resp.StatusCode != 200 {
		log.Fatal("服务器返回错误:", resp.Status)
	}

	// 创建输出CSV文件
	outputFile := strings.ReplaceAll(selectedModel, "_", " ") + ".csv"
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("创建CSV文件失败:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 读取响应并处理数据
	scanner = bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "*") {
			var fields []string
			if isCSV {
				// 如果是CSV格式，先删除所有空格，然后按逗号分割
				line = strings.ReplaceAll(line, " ", "")
				fields = strings.Split(line, ",")
				if len(fields) >= 2 {
					writer.Write([]string{fields[0], fields[1]})
				}
			} else {
				// 如果是空格分隔的格式
				fields = strings.Fields(line)
				if len(fields) >= 2 {
					writer.Write([]string{fields[0], fields[1]})
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("读取文件失败:", err)
	}

	fmt.Printf("数据已保存到CSV文件: %s\n", outputFile)
}
