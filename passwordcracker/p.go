package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type input struct {
	pass  []string
	guess string
}

var (
	inputs []input
)

func removeOne(list []string, p string) []string {
   nlist := make([]string, 0)
   for _, i := range list {
      if strings.Compare(i, p) != 0 {
         nlist = append(nlist, i)
      }
   }
   return nlist
}

func resolvePass(i string, list []string) []string {
   if len(i) == 0 {
      fmt.Printf("Found!\n")
      return []string{}
   }

   result := make([]string, 0)
   nlist := make([]string, len(list))

   copy(nlist, list)

   fmt.Printf("result: %v, nlist: %v\n", result, nlist)

   for _, p := range list {
      nlist = removeOne(nlist, p)
      n := strings.Index(i, p)

      fmt.Printf("p: %v, n: %v\n", p, n)

      if n != -1 {
         result = append(result, p)
         nstring := strings.Replace(i, p, "", -1)

         fmt.Printf("nstring: %v, nlist: %v\n", nstring, nlist)

         nl := resolvePass(nstring, nlist)
         if len(nl) > 0 {
            if strings.Compare(nl[0], "WRONG PASSWORD") == 0 {
               return nl
            } else {
               result = append(result, nl...)
            }
         }
         fmt.Printf("result: %v\n", result)
         return result
      }
   }
   return []string{"WRONG PASSWORD"}
}

func findPass(i input) []string {
	fmt.Printf("%+v\n", i)

	list := make([]string, 0)
   list = append(list, i.pass...)

   result_set := resolvePass(i.guess, list)

   fmt.Printf("result_set: %v\n", result_set)
   if len(result_set) > 0 && strings.Compare(result_set[0], "WRONG PASSWORD") == 0 {
      fmt.Printf("return WRONG PASSWORD")
      return result_set
   }

   final_list := make([]string, 0)

   for n := 0; n < len(i.guess); {
      for _, r := range result_set {
         if strings.HasPrefix(i.guess[n:], r) {
            final_list = append(final_list, r)
            n = n + len(r)
            break
         }
      }
   }

	return final_list
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTests, err := strconv.ParseInt(readLine(reader), 10, 64)

	inputs = make([]input, nTests)

	for i := 0; i < int(nTests); i++ {
		_, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		inputs[i].pass = make([]string, 0)
		np := strings.Split(readLine(reader), " ")
		inputs[i].pass = append(inputs[i].pass, np...)

		s := readLine(reader)
		inputs[i].guess = s

		result := findPass(inputs[i])
      fmt.Fprintf(writer, "%v\n", strings.Join(result, " "))
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
