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

	Package Name: day3
	Description: Subcommand for Day 3 of Advent of Code 2024
	Author: Joseph Bochinski
	Date: 2024-12-10

********************************************************************************
*/
package day3

import (
	"advent/cmn"
	"fmt"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

var mulRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var mulRe2 = regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))|(?:do\(\))|(?:don't\(\))`)

var MullItCmd = &cobra.Command{
	Use:   "mull-it",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handler := cmn.NewHandler(
			cmn.WithArgs(args),
			cmn.WithSolvers(SolvePuzzleOne, SolvePuzzleTwo),
		)
		cmn.HandleErr(handler.Solve())
	},
}

func init() {
	cmn.InitDailyCmd(MullItCmd, 3)
}

func extractMatchInts(text string) (values [][]int, err error) {
	matches := mulRe.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) == 3 {
			aStr := match[1]
			bStr := match[2]

			a, err := strconv.Atoi(aStr)
			if err != nil {
				return nil, err
			}

			b, err := strconv.Atoi(bStr)
			if err != nil {
				return nil, err
			}

			values = append(values, []int{a, b})
		}
	}

	return values, nil
}

func SolvePuzzleOne(handler *cmn.AdventHandler) error {
	defer cmn.StartProfile("SolvePuzzleOne")()

	total := 0
	for handler.Scan() {
		line := handler.Text()
		values, err := extractMatchInts(line)
		if err != nil {
			return err
		}

		for _, pair := range values {
			total += (pair[0] * pair[1])
		}
	}

	fmt.Println("Total:", total)
	return nil
}

var disabled = false

func extractMatchInts2(text string) (values [][]int, err error) {
	matches := mulRe2.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) != 3 {
			fmt.Println("bad value")
		}
		if len(match) == 3 {
			switch match[0] {
			case "don't()":
				disabled = true
			case "do()":
				disabled = false
			default:
				if disabled {
					continue
				}
				aStr := match[1]
				bStr := match[2]

				a, err := strconv.Atoi(aStr)
				if err != nil {
					fmt.Println("a:", aStr, "b:", bStr, "text:", text, "match:", match)
					return nil, err
				}

				b, err := strconv.Atoi(bStr)
				if err != nil {
					fmt.Println("a:", aStr, "b:", bStr, "text:", text)
					return nil, err
				}

				values = append(values, []int{a, b})
			}
		}
	}

	return values, nil
}

func SolvePuzzleTwo(handler *cmn.AdventHandler) error {
	defer cmn.StartProfile("SolvePuzzleTwo")()

	total := 0

	for handler.Scan() {
		line := handler.Text()
		values, err := extractMatchInts2(line)
		if err != nil {
			return err
		}

		for _, pair := range values {
			total += (pair[0] * pair[1])
		}
	}
	fmt.Println("Total:", total)

	return nil
}
