package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	nLines := strToInt(readLine())
	for id := 1; id <= nLines; id++ {
		f := strings.Split(readLine(), " ")
		w := strToInt(f[0])
		words := make([]string, w)
		for i := 0; i < w; i++ {
			words[i] = reverse(readLine())
		}

		fmt.Printf("Case #%d: %d\n", id, solve(words))
	}
}

func solve(words []string) int {
	sort.Strings(words)

	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		for k := 1; k <= len(words[i]); k++ {
			m[words[i][:k]]++
		}
	}

	keys := []string{}
	for k, v := range m {
		if v == 1 {
			delete(m, k)
		} else {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	res := 0
	for i := len(keys) - 1; i >= 0; i-- {
		if m[keys[i]] >= 2 {
			res = res + 2
			for k := range m {
				if strings.HasPrefix(keys[i], k) {
					m[k] = m[k] - 2
				}
			}
		}
	}

	return res
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

func reverse(s string) string {
	rs := []rune(s)
	out := make([]rune, len(rs))
	for i := 1; i <= len(rs); i++ {
		out[len(rs)-i] = rs[i-1]
	}
	return string(out)
}
