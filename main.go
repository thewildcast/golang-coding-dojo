package main

import (
	"fmt"
)

// Cell ...
type Cell struct {
	Top, Right, Bottom, Left bool
}

func (c *Cell) String() string {
	// 0000
	// trlb
	state := byte(0)
	if c.Top {
		state |= 0x08
	}
	if c.Right {
		state |= 0x04
	}
	if c.Bottom {
		state |= 0x02
	}
	if c.Left {
		state |= 0x01
	}

	return fmt.Sprintf("%d", state)
}

func (c *Cell) IsVisited() bool {
	return c.Top || c.Bottom || c.Left || c.Right
}

type Maze [][]*Cell

func main() {
	fmt.Println(&Cell{true,true, false, false})
	c := &Cell{false, false, false, false}
	fmt.Println(c.IsVisited())
}

// TODO: implement backtracking of maze
func backtrack(maze Maze) {}


func newMaze(width, height int) Maze {
	maze := make(Maze, height)
	for i := 0; i < height; i++ {
		maze[i] = make([]*Cell, width)
		for j := 0; j < width; j++ {
			maze[i][j] = &Cell{true, true, true, true}
		}
	}
	return maze
}