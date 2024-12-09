/*
Copyright 2024 Joseph Bochinski

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

********************************************************************************

	Package: cmn
	Title: handler
	Description: Common struct for handling Advent of Code challenges
	Author: Joseph Bochinski
	Date: 2024-12-09

********************************************************************************
*/
package cmn

import (
	"bufio"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// AdventHandler is a struct that contains common fields and methods for handling
// Advent of Code challenges
type AdventHandler struct {
	Day    string
	Puzzle string
	FileStream *os.File
	Scanner *bufio.Scanner // I may eventually add other forms of readers here
}

// NewHandler creates a pointer reference to a new AdventHandler struct and 
// initializes the filestream and reader(s)
func NewHandler(cmd *cobra.Command) (*AdventHandler, error) {
	dayNum, err := cmd.Flags().GetInt("day-num")
	if err != nil {
		dayNum = 1
	}
	puzzleNum, err := cmd.Flags().GetInt("puzzle-num")
	if err != nil {
		puzzleNum = 1
	}

	day := strconv.Itoa(dayNum)
	puzzle := strconv.Itoa(puzzleNum)
	useSample, err := cmd.Flags().GetBool("sample") 
	if err != nil {
		useSample = false
	}

	file, scanner, err := GetPuzzleDataScanner(day, puzzle, useSample)
	if err != nil {
		return nil, err
	}
	return &AdventHandler{
		Day: day,
		Puzzle: puzzle,
		FileStream: file,
		Scanner: scanner,
	}, nil
}