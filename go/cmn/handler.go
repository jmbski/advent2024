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
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

type HandlerFunc = func(handler *AdventHandler) error

// AdventHandler is a struct that contains common fields and methods for handling
// Advent of Code challenges
type AdventHandler struct {
	Day        string         // Day is a string representation of the day number, used for file paths
	DayNum     int            // DayNum is the integer representation of the day number
	Puzzle     string         // Puzzle is a string representation of the puzzle number, used for file paths
	PuzzleNum  int            // PuzzleNum is the integer representation of the puzzle number
	FileStream *os.File       // FileStream is the filestream for the puzzle data
	Scanner    *bufio.Scanner // Scanner is the bufio.Scanner for the puzzle data
	IsSample   bool           // IsSample is a boolean flag that determines if the sample data should be used

	solvers      []HandlerFunc  // solvers is a slice of functions that solve the puzzles
	cmd          *cobra.Command // cmd is a reference to the cobra command that the handler is associated with
	debugEnabled bool           // debugEnabled is a boolean flag that determines if debug output should be printed
}

// HandlerOption is a functional option type for AdventHandler
type HandlerOption func(*AdventHandler)

// WithSolvers is a functional option that allows for the assignment of solver
// functions to the AdventHandler struct
func WithSolvers(functs ...HandlerFunc) HandlerOption {
	return func(h *AdventHandler) {
		h.solvers = functs
	}
}

// WithArgs is a functional option that allows for passing CLI args to the handler
// and parse the cmd flags
func WithArgs(args []string) HandlerOption {
	return func(h *AdventHandler) {
		if h.cmd != nil && !h.cmd.Flags().Parsed() {
			HandleErr(h.cmd.Flags().Parse(args))
		}
	}
}

// NewHandler creates a pointer reference to a new AdventHandler struct and
// initializes the filestream and reader(s)
func NewHandler(opts ...HandlerOption) (h *AdventHandler) {
	if ActiveCmd == nil {
		HandleErr(&ActiveCmdUndefinedError{})
	}

	// Initialize the handler
	h = &AdventHandler{cmd: ActiveCmd}
	var err error

	// Apply the functional options
	for _, opt := range opts {
		opt(h)
	}

	// Assign/derive values from the command flags
	if h.cmd.Flags().Parsed() {
		h.DayNum = GetFlagIntD("day-num", 1)

		h.PuzzleNum = GetFlagIntD("puzzle-num", 1)

		h.Day = strconv.Itoa(h.DayNum)

		h.Puzzle = strconv.Itoa(h.PuzzleNum)

		h.IsSample = GetFlagBool("sample")

		h.debugEnabled = GetFlagBool("debug")

		if err = h.getPuzzleDataScanner(); err != nil {
			HandleErr(err)
		}
	}
	return h
}

func (h *AdventHandler) Debug(args ...any) {
	if h.debugEnabled {
		fmt.Println(args...)
	}
}

func (h *AdventHandler) Debugf(fmtStr string, args ...any) {
	if h.debugEnabled {
		fmt.Printf(fmtStr, args...)
	}
}

// Close closes the filestream
func (h *AdventHandler) Close() {
	h.FileStream.Close()
}

// Scan is a shortcut to the Scanner's Scan method
func (h *AdventHandler) Scan() bool {
	return h.Scanner.Scan()
}

// Text is a shortcut to the Scanner's Text method
func (h *AdventHandler) Text() string {
	return h.Scanner.Text()
}

// Solve Checks the command's puzzle num and if there's been a solver function
// assigned to that puzzle and executes it if so
func (h *AdventHandler) Solve() error {
	// use puzzle number -1 since the puzzle nums are not 0-based
	solverIdx := h.PuzzleNum - 1
	if solverIdx <= len(h.solvers) {
		return h.solvers[solverIdx](h)
	}
	return &SolverUndefinedError{PuzzleNum: h.PuzzleNum}
}

// getPuzzleDataScanner assigns a filestream and scanner for the puzzle data
func (h *AdventHandler) getPuzzleDataScanner() (err error) {
	fileName := "puzzle" + h.Puzzle + ".txt"
	if h.IsSample {
		fileName = "sample" + h.Puzzle + ".txt"
	}
	path := filepath.Join(DataDir, "day"+h.Day, fileName)
	h.FileStream, err = os.Open(path)
	if err != nil {
		return err
	}

	h.Scanner = bufio.NewScanner(h.FileStream)

	return nil
}
