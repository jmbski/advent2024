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

	Package Name: day1
	Description: Advent of code day 1
	Author: Joseph Bochinski
	Date: 2024-12-09

********************************************************************************
*/
package day1

import (
	"advent/cmn"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

var LocationCheck = &cobra.Command{
	Use:   "loc-check",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if err := cmd.ParseFlags(args); err != nil {
			fmt.Println(err)
			return
		}

		puzzle, err := cmd.Flags().GetInt("puzzle-num")
		if err != nil {
			fmt.Println(err)
			return
		}

		if puzzle == 1 {
			if err := SolvePuzzleOne(cmd); err != nil {
				fmt.Println(err)
			}
		} else {

			if err := SolvePuzzleTwo(cmd); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	LocationCheck.Flags().IntP("day-num", "d", 1, "The day number")
}

// SolvePuzzleOne solves the first puzzle
func SolvePuzzleOne(cmd *cobra.Command) error {
	handler, err := cmn.NewHandler(cmd)
	if err != nil {
		return err
	}
	defer handler.FileStream.Close()

	leftNums, rightNums, err := ParseP1Data(handler)
	if err != nil {
		return err
	}

	totalDist := 0
	for i, leftNum := range leftNums {
		rightNum := rightNums[i]
		dist := math.Abs(float64(leftNum - rightNum))
		totalDist += int(dist)
	}
	fmt.Println("Total distance:", totalDist)
	return nil
}

// ParseP1Data parses the data for the first puzzle
func ParseP1Data(handler *cmn.AdventHandler) (leftNums, rightNums []int, err error) {

	sepRe := regexp.MustCompile(`\s+`)
	for handler.Scanner.Scan() {
		line := handler.Scanner.Text()
		numPair := sepRe.Split(line, -1)
		if len(numPair) != 2 {
			return nil, nil, fmt.Errorf("invalid data format:\n%v\n", line)
		}
		if leftNum, err := strconv.Atoi(numPair[0]); err != nil {
			return nil, nil, err
		} else {
			leftNums = append(leftNums, leftNum)
		}

		if rightNum, err := strconv.Atoi(numPair[1]); err != nil {
			return nil, nil, err
		} else {
			rightNums = append(rightNums, rightNum)
		}
	}

	sort.Slice(leftNums, func(i, j int) bool {
		return leftNums[i] < leftNums[j]
	})

	sort.Slice(rightNums, func(i, j int) bool {
		return rightNums[i] < rightNums[j]
	})

	return leftNums, rightNums, nil
}

// SolvePuzzleTwo solves the second puzzle
func SolvePuzzleTwo(cmd *cobra.Command) error {
	handler, err := cmn.NewHandler(cmd)
	if err != nil {
		return err
	}
	defer handler.FileStream.Close()

	nums, numCounts, err := ParseP2Data(handler)
	if err != nil {
		return err
	}

	totalScore := 0
	for _, num := range nums {
		if count, exists := numCounts[num]; exists {
			totalScore += num * count
		}
	}

	fmt.Printf("Total score: %v\n", totalScore)

	return nil
}

// ParseP2Data parses the data for the second puzzle
// In this case, the goal is to check the number of times each leftNum appears
// in rightNums, multiply the number by that amount, and add the total scores
// So to parse the data right, I need to return leftNums and then a map[int]int
// with the counts per leftNum
func ParseP2Data(handler *cmn.AdventHandler) (leftNums []int, numCounts map[int]int, err error) {
	// to start, I can utilize the results from ParseP1Data still
	var rightNums []int
	leftNums, rightNums, err = ParseP1Data(handler)
	if err != nil {
		return nil, nil, err
	}

	numCounts = map[int]int{}
	for _, num := range rightNums {
		if _, exists := numCounts[num]; !exists {
			numCounts[num] = 1
		} else {
			numCounts[num]++
		}
	}
	return leftNums, numCounts, nil
}