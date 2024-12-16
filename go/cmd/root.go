/*
Copyright © 2024 Joseph Bochinski <jmbochinski@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"advent/cmd/day1"
	"advent/cmd/day10"
	"advent/cmd/day11"
	"advent/cmd/day12"
	"advent/cmd/day13"
	"advent/cmd/day14"
	"advent/cmd/day15"
	"advent/cmd/day16"
	"advent/cmd/day17"
	"advent/cmd/day18"
	"advent/cmd/day19"
	"advent/cmd/day2"
	"advent/cmd/day20"
	"advent/cmd/day21"
	"advent/cmd/day22"
	"advent/cmd/day23"
	"advent/cmd/day24"
	"advent/cmd/day25"
	"advent/cmd/day3"
	"advent/cmd/day4"
	"advent/cmd/day5"
	"advent/cmd/day6"
	"advent/cmd/day7"
	"advent/cmd/day8"
	"advent/cmd/day9"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().IntP("puzzle-num", "p", 1, "The puzzle number to run")
	rootCmd.PersistentFlags().BoolP("sample", "s", false, "Run the sample data")
	rootCmd.PersistentFlags().BoolP("debug", "D", false, "Enable debug output")
	rootCmd.AddCommand(day1.LocationCheck)
	rootCmd.AddCommand(day2.SafeReports)
	rootCmd.AddCommand(day3.MullItCmd)
	rootCmd.AddCommand(day4.WordSearch)
	rootCmd.AddCommand(day5.PrintItCmd)
	rootCmd.AddCommand(day6.Day6Cmd)
	rootCmd.AddCommand(day7.Day7Cmd)
	rootCmd.AddCommand(day8.Day8Cmd)
	rootCmd.AddCommand(day9.Day9Cmd)
	rootCmd.AddCommand(day10.Day10Cmd)
	rootCmd.AddCommand(day11.Day11Cmd)
	rootCmd.AddCommand(day12.Day12Cmd)
	rootCmd.AddCommand(day13.Day13Cmd)
	rootCmd.AddCommand(day14.Day14Cmd)
	rootCmd.AddCommand(day15.Day15Cmd)
	rootCmd.AddCommand(day16.Day16Cmd)
	rootCmd.AddCommand(day17.Day17Cmd)
	rootCmd.AddCommand(day18.Day18Cmd)
	rootCmd.AddCommand(day19.Day19Cmd)
	rootCmd.AddCommand(day20.Day20Cmd)
	rootCmd.AddCommand(day21.Day21Cmd)
	rootCmd.AddCommand(day22.Day22Cmd)
	rootCmd.AddCommand(day23.Day23Cmd)
	rootCmd.AddCommand(day24.Day24Cmd)
	rootCmd.AddCommand(day25.Day25Cmd)

}
