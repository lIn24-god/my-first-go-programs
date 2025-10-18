package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 定义颜色常量，让输出更漂亮
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorReset  = "\033[0m"
)

// 获取用户输入的函数
func getInput(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("%s无效输入，请输入数字！%s\n", ColorRed, ColorReset)
		return getInput(prompt) // 递归调用直到输入正确
	}
	return value
}

// 计算圆的面积
func calculateCircleArea(radius float64) {
	area := math.Pi * radius * radius
	fmt.Printf("%s圆的面积 = π × %.2f² = %.4f%s\n", ColorCyan, radius, area, ColorReset)
}

// 计算斐波那契数列
func calculateFibonacci(n int) {
	fmt.Printf("%s斐波那契数列前%d项：%s", ColorGreen, n, ColorReset)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}

// 判断素数
func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// 计算二次方程
func solveQuadratic(a, b, c float64) {
	discriminant := b*b - 4*a*c
	fmt.Printf("%s二次方程 %.2fx² + %.2fx + %.2f = 0 的解：%s\n", ColorPurple, a, b, c, ColorReset)

	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("两个实根：x1 = %.4f, x2 = %.4f\n", x1, x2)
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Printf("一个实根：x = %.4f\n", x)
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
		fmt.Printf("两个复根：x1 = %.4f + %.4fi, x2 = %.4f - %.4fi\n",
			realPart, imaginaryPart, realPart, imaginaryPart)
	}
}

// 生成乘法表
func generateMultiplicationTable(number float64) {
	fmt.Printf("%s%.0f的乘法表：%s\n", ColorYellow, number, ColorReset)
	for i := 1; i <= 10; i++ {
		result := number * float64(i)
		fmt.Printf("%.0f × %d = %.0f\n", number, i, result)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s🎯 欢迎使用高级数学工具包！%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s=================================%s\n", ColorBlue, ColorReset)

	for {
		fmt.Println("\n请选择要使用的功能：")
		fmt.Printf("%s1.%s 计算圆的面积\n", ColorCyan, ColorReset)
		fmt.Printf("%s2.%s 生成斐波那契数列\n", ColorGreen, ColorReset)
		fmt.Printf("%s3.%s 判断素数\n", ColorPurple, ColorReset)
		fmt.Printf("%s4.%s 解二次方程\n", ColorYellow, ColorReset)
		fmt.Printf("%s5.%s 生成乘法表\n", ColorRed, ColorReset)
		fmt.Printf("%s0.%s 退出程序\n", ColorBlue, ColorReset)
		fmt.Print("请输入选项 (0-5): ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Printf("\n%s🧮 圆的面积计算%s\n", ColorCyan, ColorReset)
			radius := getInput("请输入圆的半径: ")
			calculateCircleArea(radius)

		case "2":
			fmt.Printf("\n%s🔢 斐波那契数列生成器%s\n", ColorGreen, ColorReset)
			n := int(getInput("请输入要生成的项数: "))
			if n > 0 {
				calculateFibonacci(n)
			} else {
				fmt.Printf("%s请输入正整数！%s\n", ColorRed, ColorReset)
			}

		case "3":
			fmt.Printf("\n%s🔍 素数判断器%s\n", ColorPurple, ColorReset)
			number := int(getInput("请输入要判断的数字: "))
			if isPrime(number) {
				fmt.Printf("%s%d 是素数！%s\n", ColorGreen, number, ColorReset)
			} else {
				fmt.Printf("%s%d 不是素数%s\n", ColorRed, number, ColorReset)
			}

		case "4":
			fmt.Printf("\n%s📊 二次方程求解器%s\n", ColorYellow, ColorReset)
			a := getInput("请输入系数 a: ")
			b := getInput("请输入系数 b: ")
			c := getInput("请输入系数 c: ")
			if a == 0 {
				fmt.Printf("%sa 不能为0！%s\n", ColorRed, ColorReset)
			} else {
				solveQuadratic(a, b, c)
			}

		case "5":
			fmt.Printf("\n%s📈 乘法表生成器%s\n", ColorRed, ColorReset)
			number := getInput("请输入要生成乘法表的数字: ")
			generateMultiplicationTable(number)

		case "0":
			fmt.Printf("\n%s👋 感谢使用！再见！%s\n", ColorBlue, ColorReset)
			return

		default:
			fmt.Printf("%s❌ 无效选项，请重新选择！%s\n", ColorRed, ColorReset)
		}

		fmt.Printf("\n%s按回车键继续...%s", ColorBlue, ColorReset)
		reader.ReadString('\n')
	}
}
