/**
* @Author: huadong.hu@outlook.com
* @Date: 8/17/20 02:10
* @Desc:
 */
package main

import (
	"fmt"
	"os"
)

func main() {

	//@TODO read file name from command line
	//the argument should be `-f`
	if os.Args[1] != "-f" {
		os.Exit(-1)
	}
	sum := 0
	sum = Sum(os.Args[2])

	//DO NOT OUTPUT ANYTHING ABOVE THIS LINE
	//DO NOT MODIFY OUTPUT FORMAT!!
	fmt.Println(sum)
}
