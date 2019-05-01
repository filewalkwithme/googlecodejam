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
	d rune
}

func main() {
	nLines := strToInt(readLine())
	id := 0
	for nLines > 0 {
		id++

		f := strings.Split(readLine(), " ")
		p := strToInt(f[0])
		q := strToInt(f[1]) + 1
		people := []point{}
		for i := 0; i < p; i++ {
			f = strings.Split(readLine(), " ")
			c := strToInt(f[0])
			r := strToInt(f[1])
			//fmt.Println(f)
			d := rune(f[2][0])

			people = append(people, point{c: c, r: r, d: d})
		}

		fmt.Printf("Case #%d: %s", id, solve(q, people))
		nLines--
	}
}

func solve(q int, people []point) string {
	gridH := make([]int, q)
	gridV := make([]int, q)

	for _, p := range people {
		if p.d == 'S' {
			for i := 0; i < p.r; i++ {
				gridV[i]++
			}
		}

		if p.d == 'N' {
			for i := p.r + 1; i < q; i++ {
				gridV[i]++
			}
		}

		if p.d == 'W' {
			for i := 0; i < p.c; i++ {
				gridH[i]++
			}
		}

		if p.d == 'E' {
			for i := p.c + 1; i < q; i++ {
				gridH[i]++
			}
		}
	}

	bigH := 0
	bigV := 0

	for i := 0; i < q; i++ {
		if gridV[i] > bigV {
			bigV = gridV[i]
		}

		if gridH[i] > bigH {
			bigH = gridH[i]
		}
	}

	resH := 0
	resV := 0
	for i := 0; i < q; i++ {
		if gridH[i] == bigH {
			resH = i
			break
		}
	}

	for i := 0; i < q; i++ {
		if gridV[i] == bigV {
			resV = i
			break
		}
	}

	return fmt.Sprintf("%d %d\n", resH, resV)
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
