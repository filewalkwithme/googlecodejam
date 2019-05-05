package main

import (
	"bufio"
	"fmt"
	//"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	nLines := strToInt(readLine())
	id := 0
	for nLines > 0 {
		id++

		f := strings.Split(readLine(), " ")
		nRobots := strToInt(f[0])
		robots := make(map[string]bool)
		for i := 0; i < nRobots; i++ {
			r := readLine()
			robots[r] = true
		}

		fmt.Printf("Case #%d: %s", id, solve(robots))
		nLines--
	}
}

func solve(robots map[string]bool) string {
	res := ""

	i := 0

	// log.Println(robots)
	for {
		nIni := len(robots)

		freqs := make(map[string]int)
		for robot := range robots {

			pos := i % len(robot)

			if robot[pos] == 'R' {
				freqs["P"]++
			}

			if robot[pos] == 'S' {
				freqs["R"]++
			}

			if robot[pos] == 'P' {
				freqs["S"]++
			}
		}

		// log.Println(freqs)
		if len(freqs) == 3 {
			return "IMPOSSIBLE\n"
		}

		currMove := ""
		if len(freqs) == 2 {
			if freqs["P"] > 0 && freqs["S"] > 0 {
				currMove = "P"
			}

			if freqs["R"] > 0 && freqs["P"] > 0 {
				currMove = "R"
			}

			if freqs["S"] > 0 && freqs["R"] > 0 {
				currMove = "S"
			}
		}

		if len(freqs) == 1 {
			if freqs["P"] > 0 {
				currMove = "P"
			}

			if freqs["R"] > 0 {
				currMove = "R"
			}

			if freqs["S"] > 0 {
				currMove = "S"
			}
		}

		for robot := range robots {
			pos := i % len(robot)

			if currMove == "P" && robot[pos] == 'R' {
				delete(robots, robot)
			}

			if currMove == "R" && robot[pos] == 'S' {
				delete(robots, robot)
			}

			if currMove == "S" && robot[pos] == 'P' {
				delete(robots, robot)
			}
		}
		res = res + currMove

		nEnd := len(robots)
		if nIni == nEnd {
			return "IMPOSSIBLE\n"
		}

		if nEnd == 0 {
			break
		}
		i++
	}

	// fmt.Println(robots, maxLen)
	return fmt.Sprintf("%s\n", res)
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
