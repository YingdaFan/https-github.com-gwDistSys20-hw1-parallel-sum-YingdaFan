/**
* @Author: huadong.hu@outlook.com
* @Date: 8/17/20 02:09
* @Desc:
 */
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int = 0
	//@TODO read file name and goroutine numbers from command line and output its sum
	//the argument should be `-f` `-g`
	if os.Args[1] != "-f" || os.Args[3] != "-g" {
		os.Exit(-1)
	}
	Array, _ := strconv.Atoi(os.Args[4])
	sum = Sum(Array, os.Args[2])

	/*for runtime.NumGoroutine() > 1 {
		time.Sleep(100 * time.Millisecond)
	}*/
	//DO NOT OUTPUT ANYTHING ABOVE THIS LINE
	//DO NOT MODIFY OUTPUT FORMAT!!
	fmt.Println(sum)
}
