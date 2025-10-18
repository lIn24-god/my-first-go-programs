package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 定义一个结构体来存储爬取结果
type CrawlResult struct {
	URL   string
	Size  int64
	Error error
}

// 爬取单个URL的函数
func crawlURL(url string, results chan<- CrawlResult, wg *sync.WaitGroup) {
	defer wg.Done() // 告诉WaitGroup这个任务完成了

	start := time.Now()

	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		results <- CrawlResult{URL: url, Error: err}
		return
	}
	defer resp.Body.Close()

	// 读取响应体并计算大小
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- CrawlResult{URL: url, Error: err}
		return
	}

	elapsed := time.Since(start)

	// 发送结果到channel
	results <- CrawlResult{
		URL:   url,
		Size:  int64(len(body)),
		Error: nil,
	}

	fmt.Printf("✅ 成功爬取: %s | 大小: %d 字节 | 耗时: %v\n", url, len(body), elapsed)
}

func main() {
	fmt.Println("🚀 启动并发网络爬虫!")
	fmt.Println("======================")

	// 要爬取的URL列表
	urls := []string{
		"https://www.baidu.com",
		"https://www.qq.com",
		"https://www.taobao.com",
		"https://www.jd.com",
		"https://www.bilibili.com",
	}

	// 创建channel来接收结果
	results := make(chan CrawlResult, len(urls))

	// 使用WaitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	startTime := time.Now()

	// 为每个URL启动一个goroutine
	for _, url := range urls {
		wg.Add(1) // 增加WaitGroup的计数器
		go crawlURL(url, results, &wg)
	}

	// 启动一个goroutine来等待所有爬取任务完成
	go func() {
		wg.Wait()
		close(results) // 关闭channel，表示没有更多结果了
	}()

	// 收集并处理结果
	var totalBytes int64
	var successCount int

	fmt.Println("\n📊 爬取结果:")
	fmt.Println("----------------------")

	for result := range results {
		if result.Error != nil {
			fmt.Printf("❌ 失败: %s | 错误: %v\n", result.URL, result.Error)
		} else {
			totalBytes += result.Size
			successCount++
		}
	}

	totalTime := time.Since(startTime)

	fmt.Println("======================")
	fmt.Printf("🎉 爬取完成!\n")
	fmt.Printf("   成功爬取: %d/%d 个网站\n", successCount, len(urls))
	fmt.Printf("   总数据量: %d 字节\n", totalBytes)
	fmt.Printf("   总耗时: %v\n", totalTime)
	fmt.Printf("   Go语言的并发能力真是太强了！✨\n")
}
