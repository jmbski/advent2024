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

	Package Name: cmn
	Description: Package containing common utility functions etc... for all Advent
		of Code challenges
	Author: Joseph Bochinski
	Date: 2024-12-09

********************************************************************************
*/
package cmn

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	DataDir = "/home/joseph/coding_base/advent2024/go/data"
)

var ActiveCmd *cobra.Command

// AbsDistInt Returns the absolute value of the difference between two ints
func AbsDistInt(a, b int) int {
	dist := a - b
	if dist < 0 {
		dist *= -1
	}
	return dist
}

// InitDailyCmd is a convenience function to add the day flag to a command and set
// the ActiveCmd variable
func InitDailyCmd(cmd *cobra.Command, day int) {
	ActiveCmd = cmd
	ActiveCmd.Flags().IntP("day-num", "d", day, "Day of the Advent of Code challenge")
}

// DistInt calculates the distance between two ints and returns both the actual
// distance and the absolute value of it
func DistInt(a, b int) (dist, abs int) {
	dist = a - b
	abs = dist
	if abs < 0 {
		abs *= -1
	}
	return
}

// GetInput is a convenience function for getting terminal input from the user
func GetInput(prompt string) (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt + "\n")

	input, err = reader.ReadString('\n')
	if err != nil {

		return "", err
	}
	return strings.TrimSpace(input), nil
}

// HandleErr is a convenience function for simple top level error handling
func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// RemFromSlice removes an element from a slice at a given index
func RemFromSlice[T comparable](list []T, idx int) []T {
	newSlice := []T{}
	for i, value := range list {
		if i != idx {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}
