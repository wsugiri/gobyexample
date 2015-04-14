package main

import (
	"./util"
	"fmt"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file.Mkdir("./tmp")
	f, _ := file.CreateFile("./tmp/data.txt")
	defer f.Close()

	startTime := time.Now()
	i := 1
	seq := 0
	iterate := 500000

	fmt.Println("start iterate per:", iterate)
	for i <= 3000000 {
		f.WriteString("sample " + strconv.Itoa(i) + "\n")
		i = i + 1

		if (i % iterate) == 0 {
			seq = seq + 1

			finishTime := time.Now()
			runningTime := finishTime.Sub(startTime)
			startTime = finishTime
			fmt.Println(seq, runningTime)
		}
	}
}
