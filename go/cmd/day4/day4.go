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

	Package Name: day4
	Description: Subcommand for Day 4 of Advent of Code 2024
	Author: Joseph Bochinski
	Date: 2024-12-10

********************************************************************************
*/
package day4

import (
	"advent/cmn"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var WordSearch = &cobra.Command{
	Use:   "word-search",
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
	cmn.InitDailyCmd(WordSearch, 4)
}

type Direction int

const (
	NorthWest Direction = iota
	North
	NorthEast
	West
	East
	SouthWest
	South
	SouthEast
)

var CharMap = map[string]string{
	"X": "M",
	"M": "A",
	"A": "S",
}

var FoundWords = cmn.NewSet[string]()

type CharNode struct {
	X       int
	Y       int
	Value   string
	ID      string
	NodeMap *[][]*CharNode

	Neighbors map[Direction]*CharNode
}

func (n *CharNode) Print() {
	dirs := [][]Direction{
		{NorthWest, North, NorthEast},
		{West, -1, East},
		{SouthWest, South, SouthEast},
	}

	for _, row := range dirs {
		for _, dir := range row {
			if dir == -1 {
				fmt.Printf("[%s:'%s']", n.ID, n.Value)
			} else {
				node := n.Neighbors[dir]
				nodeStr := "[-,-:'-']"
				if node != nil {
					nodeStr = fmt.Sprintf("[%s:'%s']", node.ID, node.Value)
				}
				fmt.Printf("%s", nodeStr)
			}
		}
		fmt.Println()
	}

}

func (n *CharNode) FindNeighbors() {
	var direction Direction

	for yOffset := -1; yOffset <= 1; yOffset++ {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			if yOffset == 0 && xOffset == 0 {
				continue
			}
			x := n.X + xOffset
			y := n.Y + yOffset
			n.Neighbors[direction] = FindNode(x, y, *n.NodeMap)
			direction++
		}
	}
}

func (n *CharNode) SearchChar(direction Direction, char string, path *[]string) bool {
	if char != n.Value {
		return false
	}

	if char == "S" {
		return true
	}

	if node, ok := n.Neighbors[direction]; ok && node != nil {
		target, exists := CharMap[n.Value]
		if !exists {
			return false
		}
		*path = append(*path, n.ID)
		return node.SearchChar(direction, target, path)
	}
	return false
}

func (n *CharNode) SearchNeighbors() {
	target := CharMap[n.Value]
	for direction, node := range n.Neighbors {
		if node != nil {
			nodePath := []string{n.ID}
			if node.SearchChar(direction, target, &nodePath) {
				pathStr := strings.Join(nodePath, "_")
				FoundWords.Add(pathStr)
			}
		}
	}
}

func (n *CharNode) CheckX() bool {
	pairs := [][]*CharNode{
		{n.Neighbors[NorthWest], n.Neighbors[SouthEast]},
		{n.Neighbors[NorthEast], n.Neighbors[SouthWest]},
	}

	crossmatch := map[string]string{
		"S": "M",
		"M": "S",
	}

	for _, pair := range pairs {
		p1, p2 := pair[0], pair[1]
		if p1 == nil || p2 == nil {
			return false
		}
		if match, ok := crossmatch[p1.Value]; !ok || match != p2.Value {
			return false
		}
	}
	return true
}

func FindNode(x, y int, nodeMap [][]*CharNode) *CharNode {
	if y < 0 || x < 0 || y >= len(nodeMap) {
		return nil
	}
	for row, nodes := range nodeMap {
		if row == y {
			if x >= len(nodes) {
				return nil
			}
			return nodes[x]
		}
	}
	return nil
}

func NewNode(x, y int, value string) *CharNode {
	return &CharNode{
		X:         x,
		Y:         y,
		Value:     value,
		ID:        fmt.Sprintf("%d,%d", x, y),
		Neighbors: map[Direction]*CharNode{},
		NodeMap:   &[][]*CharNode{},
	}
}

func NewNodeMap(values [][]string, anchorChar string) (nodeMap [][]*CharNode, anchorNodes []*CharNode) {
	nodeMap = make([][]*CharNode, len(values))
	anchorNodes = []*CharNode{}
	for y, row := range values {
		nodeRow := make([]*CharNode, len(row))
		for x, value := range row {
			node := NewNode(x, y, value)
			nodeRow[x] = node
			if value == anchorChar {
				anchorNodes = append(anchorNodes, node)
			}
		}
		nodeMap[y] = nodeRow
	}

	return nodeMap, anchorNodes
}

func InitNeighbors(nodeMap [][]*CharNode) {
	//for y := 0; y < len(nodeMap); y++ {
	for _, nodes := range nodeMap {
		//nodes := nodeMap[y]
		//for x := 0; x < len(nodes); x++ {
		for _, node := range nodes {
			//node := nodes[x]
			node.NodeMap = &nodeMap
			node.FindNeighbors()
		}
	}
}

func ParseP1Data(handler *cmn.AdventHandler) []*CharNode {
	charValues := [][]string{}
	for handler.Scan() {
		line := handler.Text()
		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		charValues = append(charValues, row)
	}

	charNodes, xNodes := NewNodeMap(charValues, "X")

	InitNeighbors(charNodes)

	return xNodes
}

func SolvePuzzleOne(handler *cmn.AdventHandler) error {
	defer cmn.StartProfile("SolvePuzzleOne")()
	xNodes := ParseP1Data(handler)
	for _, node := range xNodes {
		node.SearchNeighbors()
	}

	fmt.Println("Found words:", FoundWords.Len())

	return nil
}

func ParseP2Data(handler *cmn.AdventHandler) []*CharNode {
	charValues := [][]string{}
	for handler.Scan() {
		line := handler.Text()
		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		charValues = append(charValues, row)
	}

	charNodes, xNodes := NewNodeMap(charValues, "A")

	if handler.IsSample {
		for _, row := range charValues {
			fmt.Println(row)
		}
	}

	InitNeighbors(charNodes)

	return xNodes
}

func SolvePuzzleTwo(handler *cmn.AdventHandler) error {
	defer cmn.StartProfile("SolvePuzzleOne")()

	aNodes := ParseP2Data(handler)
	fmt.Println(len(aNodes))
	count := 0
	for _, node := range aNodes {
		if node.CheckX() {
			count++
		}
	}

	fmt.Println("Found:", count)

	return nil
}
