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
	f := strings.Split(readLine(), " ")
	t := strToInt(f[0])
	n := strToInt(f[1])
	_ = strToInt(f[2])

	for i := 0; i < t; i++ {
		max := 0
		for j := 0; j < n; j++ {
			str := "18 18 18 18 18 18 18 18 18 18 18 18 18 18 18 18 18 18"

			fmt.Println(str)
			f := strings.Split(readLine(), " ")

			sum := 0
			for _, s := range f {
				sum = sum + strToInt(s)
			}
			if sum > max {
				max = sum
			}
		}

		fmt.Println(max)
		res := readLine()
		if res == "-1" {
			break
		}
	}
}

func solve(r, c int) string {
	return ""
}

// UTILS

var reader *bufio.Reader

func init() {
	rand.Seed(time.Now().UnixNano())
	reader = bufio.NewReader(os.Stdin)
}

func readLine() string {
	buf := []byte{}
	for {
		b, isPrefix, err := reader.ReadLine()
		if err != nil {
			return fmt.Sprintf("ERROR-readline: %v\n", err)
		}

		buf = append(buf, b...)
		if !isPrefix {
			break
		}

	}
	return string(buf)
}

func strToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
