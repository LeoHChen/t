package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type cell struct {
	r       uint
	c       uint
	isEmpty bool
}

type params struct {
	start    cell
	end      cell
	costWalk uint
	costJump uint
	size     uint
}

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func parseInput(filename string) (params, error) {
	var parameter params

	return parameter, nil
}

func usage() {
	fmt.Println("Usage:", os.Args[0], "inputfile")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	param, err := parseInput(os.Args[1])
	if err != nil {
		logger.Fatalf("parseInput failed: %v", err)
	}

	logger.Printf("working on maze ...")
	logger.Printf("input: %v", param)

	fmt.Print(&buf)

	os.Exit(0)
}
