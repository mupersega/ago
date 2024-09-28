package comps

import (
	"math/rand"
)

type TileMap struct {
	MaxAltitude int
	Width       int
	Height      int
	SeedData    [][]int
	Tiles       [][]Tile
}

type Tile struct {
	Id       int
	Altitude int
	X        int
	Y        int
}

type Coord struct {
	X int
	Y int
}

func NewTileMap(maxAltitude, width, height int) TileMap {
	tm := TileMap{
		MaxAltitude: maxAltitude,
		Width:       width,
		Height:      height,
		SeedData:    make([][]int, width),
	}
	for i := 0; i < width; i++ {
		tm.SeedData[i] = make([]int, height)
	}
	tm.FillSeedData()
	// tm.SeedData = tm.RandomSmooth(3)
	tm.SeedData = tm.SelectiveRandomSmooth(3, 30)
	tm.SeedData = tm.Smooth(1)
	tm.Tiles = tm.GenerateTiles()
	return tm
}

func (tm TileMap) Class() string {
	switch {
	case tm.Width == 30:
		return "small"
	case tm.Width == 50:
		return "medium"
	case tm.Width == 70:
		return "large"
	default:
		return "unknown"
	}
}

func (tm TileMap) AltAt(x, y int) int {
	return tm.SeedData[y][x]
}

func (tm *TileMap) Set(x, y, value int) {
	tm.SeedData[y][x] = value
}

func (tm *TileMap) FillSeedData() {
	for x := 0; x < tm.Width; x++ {
		for y := 0; y < tm.Height; y++ {
			tm.Set(x, y, rand.Intn(tm.MaxAltitude))
		}
	}
}

func (tm TileMap) Display() {
	print("TileMap: ")
	// print in columns and rows
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			print(tm.AltAt(x, y), " ")
		}
		println()
	}
	print(tm.Width)
}

func (tm TileMap) Smooth(distance int) [][]int {
	smoothed := make([][]int, tm.Width)
	for i := 0; i < tm.Width; i++ {
		smoothed[i] = make([]int, tm.Height)
	}
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			smoothed[y][x] = tm.SmoothPoint(x, y, distance)
		}
	}
	return smoothed
}

func (tm TileMap) SmoothPoint(x, y, distance int) int {
	total := 0
	count := 0

	for dx := -distance; dx <= distance; dx++ {
		for dy := -distance; dy <= distance; dy++ {
			wrappedX := (x + dx + tm.Width) % tm.Width
			wrappedY := (y + dy + tm.Height) % tm.Height

			total += tm.AltAt(wrappedX, wrappedY)
			count++
		}
	}

	return total / count
}

func (tm TileMap) SmoothPointsAndNeighbours(points []Coord, distance int) [][]int {
	copied := make([][]int, tm.Width)
	for i, innerSlice := range tm.SeedData {
		copied[i] = make([]int, len(innerSlice))
		copy(copied[i], innerSlice)
	}
	// get neighbours and store
	neighbours := make([]Coord, 0)
	for _, point := range points {
		neighbours = append(neighbours, tm.GetNeighbours(point.X, point.Y, distance)...)
	}
	for _, neighbour := range neighbours {
		copied[neighbour.Y][neighbour.X] = tm.SmoothPoint(neighbour.X, neighbour.Y, distance)
	}
	return copied
}

func (tm TileMap) GetNeighbours(x, y, distance int) []Coord {
	neighbours := make([]Coord, 0)
	for dx := -distance; dx <= distance; dx++ {
		for dy := -distance; dy <= distance; dy++ {
			wrappedX := (x + dx + tm.Width) % tm.Width
			wrappedY := (y + dy + tm.Height) % tm.Height
			neighbours = append(neighbours, Coord{X: wrappedX, Y: wrappedY})
		}
	}
	return neighbours
}

func (tm TileMap) RandomSmooth(maxDistance int) [][]int {
	smoothed := make([][]int, tm.Width)
	for i := 0; i < tm.Width; i++ {
		smoothed[i] = make([]int, tm.Height)
	}
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			smoothed[y][x] = tm.SmoothPoint(x, y, rand.Intn(maxDistance))
		}
	}
	return smoothed
}

func (tm TileMap) SelectiveRandomSmooth(maxDistance int, iters int) [][]int {
	// smooth pockets of tiles around the map for a number of iterations using maxDistance to determine the cluster of tiles to smooth
	smoothed := make([][]int, tm.Width)
	for i := 0; i < tm.Width; i++ {
		smoothed[i] = make([]int, tm.Height)
	}
	pointsToSmooth := make([]Coord, 0)
	for i := 0; i < iters; i++ {
		x := rand.Intn(tm.Width)
		y := rand.Intn(tm.Height)
		pointsToSmooth = append(pointsToSmooth, Coord{X: x, Y: y})
	}
	smoothed = tm.SmoothPointsAndNeighbours(pointsToSmooth, maxDistance)
	return smoothed
}

func (tm TileMap) GenerateTiles() [][]Tile {
	tiles := make([][]Tile, tm.Width)
	for i := 0; i < tm.Width; i++ {
		tiles[i] = make([]Tile, tm.Height)
	}
	i := 0
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			tiles[y][x] = Tile{
				Id:       i,
				Altitude: tm.AltAt(x, y),
				X:        x,
				Y:        y,
			}
			i++
		}
	}
	return tiles
}
