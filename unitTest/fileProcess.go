package unittest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 读取文件第一行
func readFirstLine() string {
	// 打开文件
	open, err := os.Open("testData")
	// 错误处理
	if err != nil {
		fmt.Println("Opening file failed,err:", err)
		return ""
	}
	// 基于文件流创建一个Scanner
	scanner := bufio.NewScanner(open)
	// 只读取第一行
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

// 处理文件第一行
func ProcessFirstLine() string {
	line := readFirstLine()
	destLine := strings.ReplaceAll(line, "1", "0")
	return destLine
}
