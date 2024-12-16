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

	Package Name: day2
	Description: Subcommand for Day 2 of Advent of Code 2024
	Author: Joseph Bochinski
	Date: 2024-12-10

********************************************************************************
*/
package day2

import (
	"advent/cmn"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

type Report struct {
	Increasing *bool
	Levels     []int
	FailedOnce bool
}

var DebugEnabled = false

var SafeReports = &cobra.Command{
	Use:   "safe-reports",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmn.HandleErr(cmd.ParseFlags(args))

		puzzle, err := cmd.Flags().GetInt("puzzle-num")
		cmn.HandleErr(err)

		async, err := cmd.Flags().GetBool("async")
		cmn.HandleErr(err)

		switch puzzle {
		case 1:
			if async {
				cmn.HandleErr(SolvePuzzleOneAsync())
			} else {
				cmn.HandleErr(SolvePuzzleOneSync())
			}
		case 2:
			cmn.HandleErr(SolvePuzzleTwo())
		}

	},
}

func init() {
	cmn.InitDailyCmd(SafeReports, 2)
	SafeReports.Flags().BoolP("async", "a", false, "Whether to solve using async methods")
	SafeReports.Flags().BoolVarP(&DebugEnabled, "debug", "D", false, "Enable debug statements")
}

func Debug(args ...any) {
	if DebugEnabled {
		fmt.Println(args...)
	}
}

func Debugf(fmtStr string, args ...any) {
	if DebugEnabled {
		fmt.Printf(fmtStr, args...)
	}
}

func (r *Report) CheckLevels(prev, cur *int) bool {
	Debug("Checking levels", prev, cur)
	if prev == nil || cur == nil {
		return true
	}

	change, abs := cmn.DistInt(*cur, *prev)
	Debug("change:", change, "abs:", abs)
	if abs < 1 || abs > 3 {
		return false
	}

	increasing := change > 0
	Debugf("Prev: %d, Cur: %d\nIncrease: %v, R.Increasing: %v", *prev, *cur, increasing, r.Increasing)
	if r.Increasing != nil {
		Debugf(", R.Increasing (val): %v\n", *r.Increasing)
	} else {
		Debugf("\n")
	}

	if r.Increasing == nil {
		r.Increasing = &increasing
		return true
	}

	return *r.Increasing == increasing
}

func (r *Report) IsSafe() bool {
	var prev, cur *int
	for _, level := range r.Levels {
		cur = &level
		if !r.CheckLevels(prev, cur) {
			return false
		}
		prev = &level
	}
	return true
}

func (r *Report) IsSafe2() bool {
	// 2 or less reports can't fail
	if len(r.Levels) <= 2 && !r.FailedOnce {
		return true
	}

	var cur, next *int
	i := 0
	for {
		if i >= len(r.Levels)-1 {
			break
		}

		level := r.Levels[i]
		nextLevel := r.Levels[i+1]
		Debugf("\nlevel: %d, next: %d, i: %d\n", level, nextLevel, i)

		cur = &level
		next = &nextLevel

		if !r.CheckLevels(cur, next) {
			Debug("Level failed, r.FailedOnce:", r.FailedOnce)
			if !r.FailedOnce {
				if i == len(r.Levels)-2 {
					return true
				}

				r.FailedOnce = true
				Debug("resetting values")
				for idx := range r.Levels {

					dampenedReport := &Report{
						Levels:     cmn.RemFromSlice(r.Levels, idx),
						FailedOnce: true,
					}
					if dampenedReport.IsSafe2() {
						return true
					}
				}

				/* if dampenedReport.IsSafe2() {
					Debug("Levels 1:", dampenedReport.Levels)
					return true
				} else {
					dampenedReport.Levels = cmn.RemFromSlice(r.Levels, i+1)
					dampenedReport.Increasing = nil
					Debug("Levels 2:", dampenedReport.Levels)
					return dampenedReport.IsSafe2()
				} */
			}
			return false
		}

		i++
	}
	return true
}

func NewReport(line string) (report *Report, err error) {

	report = &Report{FailedOnce: false}
	parsedLine, _, _ := strings.Cut(line, "#")
	levelStrs := strings.Split(strings.TrimSpace(parsedLine), " ")

	for _, levelStr := range levelStrs {

		if level, err := strconv.Atoi(levelStr); err != nil {
			return report, &cmn.InvalidDataError{Line: line, Err: err}
		} else {
			report.Levels = append(report.Levels, level)
		}
	}

	return report, nil
}

func SolvePuzzleOneSync() error {
	defer cmn.StartProfile("SolvePuzzleOneSync")()

	handler := cmn.NewHandler()

	defer handler.Close()
	safeCount := 0

	for handler.Scanner.Scan() {
		line := handler.Scanner.Text()
		report, err := NewReport(line)
		if err != nil {
			return err
		}

		safe := report.IsSafe()

		if safe {
			safeCount++
		}
		if handler.IsSample {
			fmt.Printf("Report: `%v` is [%v]\n", line, safe)
		}

	}

	fmt.Println("Puzzle 1 Safe count:", safeCount)
	return nil
}

func SolvePuzzleOneAsync() error {
	defer cmn.StartProfile("SolvePuzzleOneAsync")()

	handler := cmn.NewHandler()

	defer handler.Close()
	safeCount := 0

	numWorkers := 32
	var wg sync.WaitGroup

	reportChan := make(chan string, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for reportLine := range reportChan {
				report, err := NewReport(reportLine)
				cmn.HandleErr(err)

				if report.IsSafe() {
					safeCount++
				}
			}
		}()
	}

	go func() {
		for handler.Scan() {
			reportChan <- handler.Text()
		}
		close(reportChan)
	}()

	wg.Wait()

	fmt.Println("Puzzle 1 Safe count:", safeCount)
	return nil
}

func SolvePuzzleTwo() error {
	defer cmn.StartProfile("SolvePuzzleTwo")()

	handler := cmn.NewHandler()

	defer handler.Close()
	safeCount := 0

	for handler.Scanner.Scan() {
		line := handler.Scanner.Text()
		report, err := NewReport(line)
		if err != nil {
			return err
		}

		safe := report.IsSafe2()
		Debug("\nLine:", line, "\nLevels:", report.Levels, "\nSafe:", safe)
		if DebugEnabled {
			cmn.GetInput("Continue?")
		}
		if safe {
			safeCount++
		}
		if handler.IsSample {
			fmt.Printf("Report: `%v` is [%v]\n", line, safe)
		}

	}

	fmt.Println("Puzzle 2 Safe count:", safeCount)
	return nil
}
