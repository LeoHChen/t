package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the highestValuePalindrome function below.
func highestValuePalindrome(s string, n int32, k int32) string {
	var mismatch int32
	var pos = make([]int32, 0)
   var changes = make([]int32, n)

	var ns = make([]byte, n)
	var i int32

	if k >= n {
		for i := 0; i < len(s); i++ {
			ns[i] = '9'
		}
		return string(ns)
	}

	for i = 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			mismatch = mismatch + 1
			pos = append(pos, i)
		}
	}

	if mismatch > k {
		return strconv.Itoa(-1)
	}

	copy(ns, []byte(s))

	for _, i = range pos {
      if ns[i] < ns[n-i-1] {
         ns[i] = ns[n-i-1]
      } else {
         ns[n-i-1] = ns[i]
      }
      changes[i] = 1
      k = k - 1
   }

	for i = 0; k > 0 && i < n/2; i++ {
      if ns[i] != '9' {
         if (k + changes[i]) >= 2 {
            ns[i] = '9'
            ns[n-i-1] = '9'
            k = k - 2 + changes[i]
         }
      }
   }

   if k > 0 && n % 2 == 1 {
      ns[n/2] = '9'
	}

	return string(ns)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	s := readLine(reader)

	result := highestValuePalindrome(s, n, k)

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
