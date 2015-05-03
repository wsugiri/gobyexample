// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	// "bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"reflect"
	"time"
	// "os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	startTime := time.Now()
	fmt.Println("start at", startTime)

	dat, err := ioutil.ReadFile("./tmp/dat3.txt")
	check(err)
	fmt.Println("type of dat", reflect.TypeOf(dat), string(dat))

	finishTime := time.Now()
	fmt.Println("finish at", finishTime)

	runningTime := finishTime.Sub(startTime)

	fmt.Println("running time", runningTime.String())
	fmt.Println("running per 1000 line", (runningTime / 5000).String())
}
