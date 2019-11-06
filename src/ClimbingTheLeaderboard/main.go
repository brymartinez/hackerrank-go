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

func removeDuplicates(a []int) []int {
	r := []int{}
	seen := map[int]bool{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			r = append(r, val)
			seen[val] = true
		}
	}
	return r
}

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int, alice []int) []int {
	scores = removeDuplicates(scores)
	fmt.Println(scores)
	rank := make([]int, len(alice))
	for j, v := range alice {
		i := sort.Search(len(scores), func(i int) bool { return scores[i] <= v })
		rank[j] = i + 1
		// if i < len(scores) && scores[i] == v {
		//     fmt.Println("entered")
		//     rank[j] = i + 1
		// }
	}
	fmt.Println(rank)
	return rank
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*3)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024*3)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
