package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the balancedSums function below.
func balancedSums(arr []int) string {
	sumLeft := 0
	sumRight := 0

	for i := range arr {
		sumRight += arr[i]
	}

	if len(arr) == 1 {
		return "YES"
	} // Catch single array entry

	for _, v := range arr {
		sumRight -= v
		if sumRight == sumLeft {
			return "YES"
		}
		sumLeft += v
	}
	return "NO"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024*3)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024*3)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int(nTemp)

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := balancedSums(arr)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
