package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type cellType struct {
	r       int
	c       int
	isEmpty bool
	minCost uint
}

type inputParams struct {
	start    cellType
	end      cellType
	costWalk uint
	costJump uint
	size     uint
}

var (
	step  = 0
	debug = false
	maze  [][]cellType
)

func parseInput(filename string) (inputParams, error) {
	var parameter inputParams

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return parameter, fmt.Errorf("unable to read file: %v", filename)
	}

	lines := strings.Split(string(content), "\n")

	for num, line := range lines {
		//      fmt.Printf("line: %v\n", line)
		if num < 7 {
			uintVal, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				return parameter, fmt.Errorf("unable to parse uint: %v", line)
			}
			switch num {
			case 0:
				parameter.start.r = int(uintVal)
			case 1:
				parameter.start.c = int(uintVal)
			case 2:
				parameter.end.r = int(uintVal)
			case 3:
				parameter.end.c = int(uintVal)
			case 4:
				parameter.costWalk = uint(uintVal)
			case 5:
				parameter.costJump = uint(uintVal)
			case 6:
				parameter.size = uint(uintVal)
				maze = make([][]cellType, parameter.size)
				for i := range maze {
					maze[i] = make([]cellType, parameter.size)
				}
			}
		} else {
			i := num - 7
			chars := strings.Split(line, "")
			for j := range chars {
				if chars[j] == "." {
					maze[i][j] = cellType{i, j, true, 0}
				}
				if chars[j] == "#" {
					maze[i][j] = cellType{i, j, false, 0}
				}
			}
		}
	}

	if debug {
		fmt.Printf("input: %+v\n", parameter)
		fmt.Printf("maze: %+v\n", maze)
	}
	return parameter, nil
}

func popMinCostCell(q []*cellType) ([]*cellType, *cellType) {
	if len(q) == 0 {
		return nil, nil
	}
	minCost := q[0].minCost
	var small *cellType = q[0]

	if len(q) == 1 {
		nq := make([]*cellType, 0)
		return nq, small
	}

	var n int
	for i := 1; i < len(q); i++ {
		if minCost > q[i].minCost {
			minCost = q[i].minCost
			small = q[i]
			n = i
		}
	}
	nq := make([]*cellType, len(q)-1)
	nq = append(q[:n], q[(n+1):]...)

	return nq, small
}

func findNextCells(q *cellType, param inputParams) ([]*cellType, []uint) {
	cells := make([]*cellType, 0)
	costs := make([]uint, 0)

	// north
	for r := q.r - 1; r >= 0; r-- {
		if debug {
			fmt.Printf("checking north cell => maze[%v][%v]\n", r, q.c)
		}
		if maze[r][q.c].isEmpty {
			cells = append(cells, &maze[r][q.c])
			if r == q.r-1 {
				costs = append(costs, param.costWalk)
			} else {
				costs = append(costs, param.costJump)
			}
			break
		}
	}
	// south
	for r := q.r + 1; r < int(param.size); r++ {
		if debug {
			fmt.Printf("checking south cell => maze[%v][%v]\n", r, q.c)
		}
		if maze[r][q.c].isEmpty {
			cells = append(cells, &maze[r][q.c])
			if r == q.r+1 {
				costs = append(costs, param.costWalk)
			} else {
				costs = append(costs, param.costJump)
			}
			break
		}
	}

	// west
	for c := q.c - 1; c >= 0; c-- {
		if debug {
			fmt.Printf("checking west cell => maze[%v][%v]\n", q.r, c)
		}
		if maze[q.r][c].isEmpty {
			cells = append(cells, &maze[q.r][c])
			if c == q.c-1 {
				costs = append(costs, param.costWalk)
			} else {
				costs = append(costs, param.costJump)
			}
			break
		}
	}
	// east
	for c := q.c + 1; c < int(param.size); c++ {
		if debug {
			fmt.Printf("checking east cell => maze[%v][%v]\n", q.r, c)
		}
		if maze[q.r][c].isEmpty {
			cells = append(cells, &maze[q.r][c])
			if c == q.c+1 {
				costs = append(costs, param.costWalk)
			} else {
				costs = append(costs, param.costJump)
			}
			break
		}
	}

	return cells, costs
}

func findMinCost(q []*cellType, param inputParams) int {
	if debug {
		fmt.Printf("step: %v\n", step)
	}

	step = step + 1

	nq, q1 := popMinCostCell(q)
	if q1 == nil {
		return -1
	}
	if debug {
		spew.Printf("popMinCostCell nq: %+v\n", nq)
		spew.Printf("popMinCostCell small: %+v\n", q1)
	}
	if q1.r == param.end.r && q1.c == param.end.c {
		return int(q1.minCost)
	}
	nextCells, costs := findNextCells(q1, param)
	for n := 0; n < len(nextCells); n++ {
		if debug {
			fmt.Printf("checking next: %+v, %v\n", nextCells[n], costs[n])
		}
		inQ := false
		for i := 0; i < len(nq); i++ {
			if nq[i] == nextCells[n] {
				inQ = true
				if costs[n]+q1.minCost < nq[i].minCost {
					nq[i].minCost = costs[n] + q1.minCost
					break
				}
			}
		}
		if !inQ {
			nextCells[n].minCost = costs[n] + q1.minCost
			nq = append(nq, nextCells[n])
		}
	}
	return findMinCost(nq, param)
}

func resolver(filename string) int {
	param, err := parseInput(filename)
	if err != nil {
		fmt.Printf("parseInput failed: %v", err)
		os.Exit(1)
	}
	if debug {
		fmt.Printf("input: %+v\n", param)
	}

	if !maze[param.start.r][param.start.c].isEmpty || !maze[param.end.r][param.end.c].isEmpty {
		return -1
	}

	q := make([]*cellType, 1)
	q[0] = &maze[param.start.r][param.start.c]

	cost := findMinCost(q, param)

	step = 0

	return cost
}

func main() {
	debugPtr := flag.Bool("debug", false, "verbose output for debugging")
	input := flag.String("input", "testfile1", "input test filename")

	flag.Parse()

	debug = *debugPtr

	cost := resolver(*input)
	fmt.Println("minCost:", cost)

	os.Exit(0)
}
