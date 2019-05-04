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
	l := strings.Split(readLine(), " ")
	t := strToInt(l[0])
	// f := strToInt(l[1])

	for i := 0; i < t; i++ {
		possible := map[string]bool{
			"A": true,
			"B": true,
			"C": true,
			"D": true,
			"E": true,
		}

		mFreq := make(map[string]int)
		mPosi := make(map[string][]int)
		firstChar := ""
		for j := 0; j < 119; j++ {
			pos := j*5 + 1
			fmt.Println(pos)
			s := readLine()
			mPosi[s] = append(mPosi[s], pos)
			mFreq[s]++
		}

		for char, freq := range mFreq {
			if freq == 23 {
				firstChar = char
				break
			}
		}

		//log.Printf("%+v", mFreq)
		//log.Printf("%+v", mPosi)

		mFreq2 := make(map[string]int)
		mPosi2 := make(map[string][]int)
		secondChar := ""
		for j := 0; j < 23; j++ {
			pos := mPosi[firstChar][j] + 1
			fmt.Println(pos)
			s := readLine()
			mPosi2[s] = append(mPosi2[s], pos)
			mFreq2[s]++
		}

		for char, freq := range mFreq2 {
			if freq == 5 {
				secondChar = char
				break
			}
		}

		//log.Printf("%+v", mFreq2)
		//log.Printf("%+v", mPosi2)

		mFreq3 := make(map[string]int)
		mPosi3 := make(map[string][]int)
		thirdChar := ""
		for j := 0; j < 5; j++ {
			pos := mPosi2[secondChar][j] + 1
			fmt.Println(pos)
			s := readLine()
			mPosi3[s] = append(mPosi3[s], pos)
			mFreq3[s]++
		}

		for char, freq := range mFreq3 {
			if freq == 1 {
				thirdChar = char
				break
			}
		}

		//log.Printf("%+v", mFreq3)
		//log.Printf("%+v", mPosi3)

		fmt.Println(mPosi3[thirdChar][0] + 1)
		element := readLine()

		guess := firstChar + secondChar + thirdChar

		delete(possible, firstChar)
		delete(possible, secondChar)
		delete(possible, thirdChar)
		delete(possible, element)

		for k := range possible {
			guess = guess + string(k)
		}
		guess = guess + element
		//log.Printf("guess: %+v", guess)

		fmt.Println(guess)
		res := readLine()
		if res == "N" {
			break
		}
	}
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
