/**
* @Author: huadong.hu@outlook.com
* @Date: 8/10/20 10:20
* @Desc:
 */
package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

// Read a list of integers from `fileName`
// and launch `goRoutineNums` go routines to do calculation
// return sum of these Integers
// You Must start exact `goRoutineNums` go routines or you lose points here

var sum int = 0

func Sum(goRoutineNums int, fileName string) int {
	//TODO Add your code here
	A, err := readInts(fileName)
	if err != nil {
		log.Fatal(err)
	}

	seg1 := float64(len(A)) / float64(goRoutineNums)
	seg := int(math.Floor(seg1))
	ch := make(chan int)

	for i := 0; i < goRoutineNums; i++ {

		l := i * seg
		h := (i+1)*seg - 1
		if i == goRoutineNums-1 {
			h = len(A) - 1
		}

		go func(A []int, l int, h int) {
			var sum1 int = 0
			for j := l; j <= h; j++ {
				sum1 += A[j]
			}
			ch <- sum1
		}(A, l, h)

	}

	chSum := 0
	counter := 1
	for v := range ch {
		chSum += v
		if counter == goRoutineNums {
			close(ch)
		}
		counter++
	}
	return chSum

}

//Read integers from file
//Do not modify this function
func readInts(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cin := bufio.NewScanner(file)
	cin.Split(bufio.ScanWords)
	var res []int
	for cin.Scan() {
		val, err := strconv.Atoi(cin.Text())
		if err != nil {
			return res, err
		}
		res = append(res, val)
	}
	return res, nil
}
