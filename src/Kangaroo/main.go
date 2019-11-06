package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the kangaroo function below.
func kangaroo(x1 int, v1 int, x2 int, v2 int) string {
	if x2 > x1 && v2 > v1 { // Check for imaginary
		return "NO"
	} else if x1 != x2 && v2 == v1 { // Check for zero divisor
		return "NO"
	} else {
		y := (x1 - x2) % (v2 - v1) // modulo, if > 0 means there's a decimal in the # of steps
		if y == 0 {
			return "YES"
		} else {
			return "NO"
		}
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*3)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024*3)

	x1V1X2V2 := strings.Split(readLine(reader), " ")

	x1Temp, err := strconv.ParseInt(x1V1X2V2[0], 10, 64)
	checkError(err)
	x1 := int(x1Temp)

	v1Temp, err := strconv.ParseInt(x1V1X2V2[1], 10, 64)
	checkError(err)
	v1 := int(v1Temp)

	x2Temp, err := strconv.ParseInt(x1V1X2V2[2], 10, 64)
	checkError(err)
	x2 := int(x2Temp)

	v2Temp, err := strconv.ParseInt(x1V1X2V2[3], 10, 64)
	checkError(err)
	v2 := int(v2Temp)

	result := kangaroo(x1, v1, x2, v2)

	fmt.Fprintf(writer, "%s\n", result)

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
