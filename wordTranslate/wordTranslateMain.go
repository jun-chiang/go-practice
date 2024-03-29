package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jun-chiang/go-practice/wordTranslate/domain"
)

// 发送封装好的http请求，并返回Response
func sendHttp(body io.Reader) (*http.Response, error) {
	// 新建客户端
	client := &http.Client{}
	// 新建请求对象
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", body)
	// 错误处理
	if err != nil {
		log.Fatal(err)
	}
	// 设置一堆参数
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("app-name", "xy")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "34fb7922cfdb90bab1c2676d847f1ece")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	// 发起http请求
	resp, err := client.Do(req)
	return resp, err
}

// 翻译单词，返回自定义响应对象
func translate(word string) domain.DictResponse {
	// 创建自定义的请求对象
	dictRequest := domain.DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(dictRequest)
	if err != nil {
		log.Fatal(err)
	}
	// 从字节数组创建 *bytes.Reader
	data := bytes.NewReader(buf)
	// 调用自定义的请求方法,获取response
	resp, err := sendHttp(data)
	// 异常处理
	if err != nil {
		log.Fatal(err)
	}
	// 读取response中的数据
	bodyText, err := io.ReadAll(resp.Body)
	// 延迟关闭流
	defer resp.Body.Close()
	// 异常处理
	if err != nil {
		log.Fatal(err)
	}
	// 反序列化成对象
	dictResponse := domain.DictResponse{}
	err = json.Unmarshal(bodyText, &dictResponse)
	// 异常处理
	if err != nil {
		log.Fatal(err)
	}
	return dictResponse
}

func main() {
	// 基于标准输入流创建一个缓冲读取器
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入要查询的单词,输入#退出程序:")
		// 从控制台获取输入
		input, err := reader.ReadString('\n')
		// 错误处理
		if err != nil {
			log.Fatal(err)
		}
		// 去除多余的\r\n
		input = strings.TrimSuffix(input, "\r\n")
		// 判断是否要退出程序
		if input == "#" {
			fmt.Println("期待您的下次使用，再见！")
			break
		}
		// 调用封装好的translate方法
		result := translate(input)
		dictionary := &result.Dictionary
		fmt.Println()
		// 打印单词以及音标
		fmt.Println(dictionary.Entry, dictionary.Prons.En)
		// 循环打印单词的释义
		for _, explanation := range dictionary.Explanations {
			fmt.Println(explanation)
		}
		fmt.Println()
	}

}
