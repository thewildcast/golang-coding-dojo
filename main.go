package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Cell ...
type Cell struct {
	Top, Right, Bottom, Left bool
	X, Y                     int
}

func (c *Cell) DropWall(vecina *Cell) {
	if c.X == vecina.X {
		if c.Y > vecina.Y {
			c.Top = false
			vecina.Bottom = false
		}
		if c.Y < vecina.Y {
			c.Bottom = false
			vecina.Top = false
		}
		return
	}

	if c.X > vecina.X {
		c.Left = false
		vecina.Right = false
	}
	if c.X < vecina.X {
		c.Right = false
		vecina.Left = false
	}
}

func (c *Cell) IsNotVisited() bool {
	return c.Top || c.Bottom || c.Left || c.Right
}

type Maze [][]*Cell

func (m Maze) Render(renderer *sdl.Renderer) {
	height := int32(800 / len(m))
	width := int32(800 / len(m[0]))

	for y, cells := range m {
		i := int32(y)
		for x, cell := range cells {
			j := int32(x)

			if cell.Top {
				renderer.DrawLine(j*width, i*height, j*width+width, i*height)
			}

			if cell.Bottom {
				renderer.DrawLine(j*width, i*height+height, j*width+width, i*height+height)
			}

			if cell.Left {
				renderer.DrawLine(j*width, i*height, j*width, i*height+height)
			}

			if cell.Right {
				renderer.DrawLine(j*width+width, i*height, j*width+width, i*height+height)
			}
		}
	}
}

func (m Maze) getNonVisitedCell(current *Cell) (*Cell, error) {
	var possibleCells []*Cell

	if current.Y-1 >= 0 && m[current.Y-1][current.X].IsNotVisited() {
		possibleCells = append(possibleCells, m[current.Y-1][current.X])
	}
	if current.Y+1 < len(m) && m[current.Y+1][current.X].IsNotVisited() {
		possibleCells = append(possibleCells, m[current.Y+1][current.X])
	}
	if current.X-1 >= 0 && m[current.Y][current.X-1].IsNotVisited() {
		possibleCells = append(possibleCells, m[current.Y][current.X-1])
	}
	if current.X+1 < len(m[0]) && m[current.Y][current.X+1].IsNotVisited() {
		possibleCells = append(possibleCells, m[current.Y][current.X+1])
	}

	if len(possibleCells) == 0 {
		return nil, fmt.Errorf("Backgrackiiiiing")
	}
	n := rand.Intn(len(possibleCells))

	return possibleCells[n], nil
}

func (m Maze) Build() {
	currentCell := m[0][0]
	var visited []*Cell
	visited = append(visited, currentCell)

	for len(visited) > 0 {
		fmt.Printf("%+v\n", currentCell)
		nextCell, err := m.getNonVisitedCell(currentCell)
		if err != nil {
			// no hay celdas alrededor sin visitar
			fmt.Println("fail")
			currentCell = visited[len(visited)-1]
			visited = visited[:len(visited)-1]
			continue
		}

		currentCell.DropWall(nextCell)
		visited = append(visited, currentCell)
		currentCell = nextCell
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(1)
	}
	defer renderer.Destroy()

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()

	maze := newMaze(10, 10)
	maze.Build()
	renderer.SetDrawColor(0, 0, 0, 255)
	maze.Render(renderer)

	renderer.Present()
	sdl.PollEvent()
	sdl.Delay(20000)
}

// TODO: implement backtracking of maze
func backtrack(maze Maze) {}

func newMaze(width, height int) Maze {
	maze := make(Maze, height)
	for i := 0; i < height; i++ {
		maze[i] = make([]*Cell, width)
		for j := 0; j < width; j++ {
			maze[i][j] = &Cell{true, true, true, true, j, i}
		}
	}
	return maze
}
