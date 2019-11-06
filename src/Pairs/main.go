package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the pairs function below.
func pairs(k int, arr []int) int {
	pairs := 0
	sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })
	// Better to sort descending because the left side can be ignored after
	// You can quit once you find the target per iteration
	// Coz there's only one solution to x - y = k

	for i, x := range arr {
		for j := i + 1; j < len(arr); j++ {
			// arr[j] will always be less than x
			if x-arr[j] == k {
				pairs++
				break
			}
		}
	}

	return pairs
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*3)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024*3)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int(kTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
