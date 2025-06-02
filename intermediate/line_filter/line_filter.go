package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var file, err = os.Open("./intermediate/line_filter/cloud.txt")
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	var lineNo = 1

	defer func() {
		file.Close()
		fmt.Println("file closed")
	}()

	var newBuffer = bufio.NewScanner(file)

	var keyword = "Cloud"

	for newBuffer.Scan() {
		line := newBuffer.Text()

		if strings.Contains(line, keyword) {
			wordReplacedString := strings.ReplaceAll(line, keyword, "Badal")
			fmt.Println(lineNo, "filtered lines:", wordReplacedString)
		}
		lineNo++
	}

	err = newBuffer.Err()

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

}
