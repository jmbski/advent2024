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

	Package Name: day5
	Description: Subcommand for Day 5 of Advent of Code 2024
	Author: Joseph Bochinski
	Date: 2024-12-10

********************************************************************************
*/
package day5

import (
	"advent/cmn"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var PrintItCmd = &cobra.Command{
	Use:   "print-it",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handler := cmn.NewHandler(
			cmn.WithSolvers(SolvePuzzleOne, SolvePuzzleTwo),
			cmn.WithArgs(args),
		)
		cmn.HandleErr(handler.Solve())
	},
}

func init() {
	cmn.InitDailyCmd(PrintItCmd, 5)
}

type OrderCheckerP1 struct {
	CheckedPages *cmn.Set[string]
	OrderMap     map[string][]string

	Scan   func() bool
	Text   func() string
	Debug  func(args ...any)
	Debugf func(str string, args ...any)
}

func (o *OrderCheckerP1) CheckPage(page string) bool {
	o.Debug("\nChecking page:", page)
	if afterPages, exists := o.OrderMap[page]; exists {
		o.Debug("After pages:", afterPages)
		o.Debug("Checked Pages:", o.CheckedPages.Values())
		for _, afterPage := range afterPages {
			if o.CheckedPages.Contains(afterPage) {
				o.Debugf("Page `%s` failed\n", page)
				return false
			}
		}
	}
	o.CheckedPages.Add(page)
	o.Debugf("Page `%s` passed\n", page)
	return true
}

func (o *OrderCheckerP1) CheckManual(pages []string) bool {
	o.CheckedPages = cmn.NewSet[string]()
	for _, page := range pages {
		if !o.CheckPage(page) {
			return false
		}
	}
	return true
}

func (o *OrderCheckerP1) InspectManuals() int {
	total := 0
	for o.Scan() {
		line := o.Text()
		pages := strings.Split(line, ",")
		if o.CheckManual(pages) {
			o.Debug("Manual Passed:", pages)
			midIdx := (len(pages) / 2)
			middle, err := strconv.Atoi(pages[midIdx])
			cmn.HandleErr(err)
			o.Debugf("Len: %d, midIdx: %d, midPage: %d\n", len(pages), midIdx, middle)
			total += middle
		}
	}

	return total
}

func NewOrderCheckerP1(h *cmn.AdventHandler) *OrderCheckerP1 {
	o := &OrderCheckerP1{
		CheckedPages: cmn.NewSet[string](),
		OrderMap:     map[string][]string{},

		Scan:   h.Scan,
		Text:   h.Text,
		Debug:  h.Debug,
		Debugf: h.Debugf,
	}

	for h.Scan() {
		line := h.Text()
		if line == "" {
			break
		}
		pageX, pageY, ok := strings.Cut(line, "|")
		if !ok {
			cmn.HandleErr(&cmn.InvalidDataError{Line: line})
		}
		o.OrderMap[pageX] = append(o.OrderMap[pageX], pageY)
	}

	return o
}

func SolvePuzzleOne(handler *cmn.AdventHandler) error {
	defer cmn.StartProfile("SolvePuzzleOne")()

	orderChecker := NewOrderCheckerP1(handler)

	if handler.IsSample {
		for key, value := range orderChecker.OrderMap {
			fmt.Printf("[%s]: %v\n", key, value)
		}
	}

	goodManuals := orderChecker.InspectManuals()

	fmt.Println("Good manual score:", goodManuals)

	return nil
}

func SolvePuzzleTwo(handler *cmn.AdventHandler) error {
	defer cmn.StartProfile("SolvePuzzleTwo")()

	return nil
}
