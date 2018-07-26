package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type cellType struct {
	r       uint
	c       uint
	isEmpty bool
}

type inputParams struct {
	start    cellType
	end      cellType
	costWalk uint
	costJump uint
	size     uint
}

var (
	debug  = true
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func parseInput(filename string) (inputParams, error) {
	var parameter inputParams

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return parameter, fmt.Errorf("unable to read file: %v", filename)
	}

	lines := strings.Split(string(content), "\n")

	for num, line := range lines {
		if num < 7 {
			uintVal, err := strconv.ParseUint(line, 10, 32)
			if err != nil {
				return parameter, fmt.Errorf("unable to parse uint: %v", line)
			}
			switch num {
			case 0:
				parameter.start.r = uint(uintVal)
			case 1:
				parameter.start.c = uint(uintVal)
			}
		}
	}

	return parameter, nil
}

func usage() {
	fmt.Println("Usage:", os.Args[0], "inputfile")
	os.Exit(1)
}

func resolver(filename string) int {
	param, err := parseInput(filename)
	if err != nil {
		logger.Fatalf("parseInput failed: %v", err)
	} else {
		logger.Printf("input: %v", param)
	}
	return 4
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	logger.Printf("working on maze ...")

	resolver(os.Args[1])

	if debug {
		fmt.Print(&buf)
	}

	os.Exit(0)
}
