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

type point struct {
	r int
	c int
}

func main() {
	nLines := strToInt(readLine())
	for id := 1; id <= nLines; id++ {
		f := strings.Split(readLine(), " ")
		r := strToInt(f[0])
		c := strToInt(f[1])

		fmt.Printf("Case #%d: %s", id, solve(r, c))
	}
}

func solve(r, c int) string {
	if r*c <= 9 {
		return "IMPOSSIBLE\n"
	}

	valid := true
	var moves []point

	for {
		valid = true

		possible := []point{}
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				possible = append(possible, point{r: i, c: j})
			}
		}

		moves = []point{}
		for t := 1; t <= r*c; t++ {
			i := rand.Intn(len(possible))
			if t > 1 {
				found := false
				for k := 0; k < 5; k++ {
					src := moves[t-2]
					dst := possible[i]

					if !(src.r == dst.r ||
						src.c == dst.c ||
						src.r-src.c == dst.r-dst.c ||
						src.r+src.c == dst.r+dst.c) {
						found = true
						break
					}
					i = rand.Intn(len(possible))
				}
				if !found {
					valid = false
					break
				}
			}

			moves = append(moves, possible[i])
			possible2 := possible[i+1:]
			possible = possible[:i]
			possible = append(possible, possible2...)
		}

		if valid {
			break
		}
	}

	if valid {
		res := "POSSIBLE\n"
		for _, v := range moves {
			res = res + fmt.Sprintf("%d %d\n", v.r+1, v.c+1)
		}

		return res
	}

	return "IMPOSSIBLE\n"
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
