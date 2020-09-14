/**
* @Author: huadong.hu@outlook.com
* @Date: 7/7/20 9:44 PM
* @Desc:
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"time"
)

var sumCli = SumServiceClient{}

func SumHandler(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte(err.Error()))
		return
	}
	var fileName string
	var goRoutineNums int
	var totalSum int
	//TODO You should add your code here
	//HINT: Receive fileName(f) and goRoutineNums(g) from URL
	fileName = r.FormValue("f")
	goRoutineNums, _ = strconv.Atoi(r.FormValue("g"))
	if fileName == "" {
		return
	}

	//HINT: Call sumCli.Sum(fileName, goRoutineNums) to retrieve sum result
	totalSum, _ = sumCli.Sum(fileName, goRoutineNums)
	//DO NOT MODIFY OUTPUT FORMAT!
	type OutPut struct {
		FileName      string `json:"file_name"`
		GoRoutineNums int    `json:"go_routine_nums"`
		TotalSum      int    `json:"sum"`
	}

	out := OutPut{
		FileName:      fileName,
		GoRoutineNums: goRoutineNums,
		TotalSum:      totalSum,
	}
	data, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	wr.Write(data)
	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

func main() {
	time.Sleep(100 * time.Millisecond)
	go regRpcService()
	handleHttp()
}

func regRpcService() {
	client, err := rpc.Dial("tcp", ":8083")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	sumCli.Client = client
}

func handleHttp() {
	http.HandleFunc("/", SumHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
