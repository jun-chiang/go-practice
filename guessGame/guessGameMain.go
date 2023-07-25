package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 随机数的最大值约束
	maxNum := 100
	// 初始化种子
	rand.Seed(time.Now().UnixNano())
	// 生成随机数
	guessResult := rand.Intn(maxNum)
	// 获取一个缓冲读取器
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please input your guess:")
		// 读取用户输入，以换行符结束
		userGuessStr, err := reader.ReadString('\n')
		// 错误处理
		if err != nil {
			fmt.Println("An error occurred while reading input.Please try again. ", err)
			continue
		}
		// 截取掉最后的换行符
		userGuessStr = strings.TrimSuffix(userGuessStr, "\r\n")
		// 把字符串转换成数字
		userGuess, err := strconv.Atoi(userGuessStr)
		// 字符串转数字的错误处理
		if err != nil {
			fmt.Println("An error occurred while parse your input number.Please try again. ", err)
			continue
		}
		// 判断用户是否猜测正确
		if userGuess > guessResult {
			fmt.Printf("Your guess %d is bigger than the actual number", userGuess)
		} else if userGuess < guessResult {
			fmt.Printf("Your guess %d is smaller than the actual number", userGuess)
		} else {
			fmt.Printf("Your guess %d is correct!", userGuess)
			return
		}
	}
}
