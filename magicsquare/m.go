package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

type matrix struct {
   m [3][3]int32
}

func findDifference(s [][]int32, sol [3][3]int32) int32 {
   var d1, d2, d3, d4 float64
   for i := range s {
      for j := range s[i] {
         d1 = d1 + math.Abs(float64(sol[i][j] - s[i][j]))
         d2 = d2 + math.Abs(float64(sol[i][2-j] - s[i][j]))
         d3 = d3 + math.Abs(float64(sol[2-i][j] - s[i][j]))
         d4 = d4 + math.Abs(float64(sol[2-i][2-j] - s[i][j]))
      }
   }
   fmt.Printf("d1:%v, d2:%v, d3:%v, d4:%v\n", d1, d2, d3, d4)
   return int32(math.Min(math.Min(d1, d2), math.Min(d3, d4)))
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	var solutions = make([]matrix, 4)
   solutions[0] = matrix { [3][3]int32{{8,1,6},
                           {3,5,7},
                           {4,9,2}} }
   solutions[1] = matrix { [3][3]int32{{4,3,8},
                           {9,5,1},
                           {2,7,6}} }
   solutions[2] = matrix { [3][3]int32{{2,9,4},
                           {7,5,3},
                           {6,1,8}} }
   solutions[3] = matrix { [3][3]int32{{6,7,2},
                           {1,5,9},
                           {8,3,4}} }

   d1 := int32(100)
   for _, sol := range solutions {
      d2 := findDifference(s, sol.m)
      if d2 < d1 {
         d1 = d2
      }
   }
   return d1
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var s [][]int32
    for i := 0; i < 3; i++ {
        sRowTemp := strings.Split(readLine(reader), " ")

        var sRow []int32
        for _, sRowItem := range sRowTemp {
            sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
            checkError(err)
            sItem := int32(sItemTemp)
            sRow = append(sRow, sItem)
        }

        if len(sRow) != 3 {
            panic("Bad input")
        }

        s = append(s, sRow)
    }

    result := formingMagicSquare(s)

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

