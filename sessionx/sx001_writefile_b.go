// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	createFile1()
	createFile2()
	createFile3()
	readFile()
}

func createFile1() {
	f, err := os.Create("./tmp/dat3.txt")
	check(err)
	defer f.Close()

	startTime := time.Now()

	i := 1
	for i <= 1000000 {
		f.WriteString("sample " + strconv.Itoa(i) + "\n")
		i = i + 1
	}

	finishTime := time.Now()
	// fmt.Println("finish at", finishTime)

	runningTime := finishTime.Sub(startTime)

	fmt.Println("running time (1)", runningTime.String())
}

func createFile2() {
	f, err := os.Create("./tmp/dat4.txt")
	check(err)
	defer f.Close()

	startTime := time.Now()

	i := 1
	for i <= 500000 {
		f.WriteString("sample " + string(i) + "\n")
		i = i + 1
	}

	finishTime := time.Now()
	runningTime := finishTime.Sub(startTime)

	fmt.Println("running time (2)", runningTime.String())
}

func createFile3() {
	f, err := os.Create("./tmp/dat5.txt")
	check(err)
	defer f.Close()

	i := 1
	for i <= 10 {
		f.WriteString("sample " + strconv.Itoa(i) + "\n")
		i = i + 1
	}
}

func readFile() {
	data, err := ioutil.ReadFile("./tmp/dat5.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	fmt.Println(strings.Split(lines[3], " ")[0])
	fmt.Println(strings.Split(lines[3], " ")[1])
}
