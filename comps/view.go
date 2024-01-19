package comps

import "math/rand"

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
	tm.SeedData = tm.Smooth(2)
	tm.SeedData = tm.RandomSmooth()
	tm.Tiles = tm.GenerateTiles()
	return tm
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
	for x := 0; x < tm.Width; x++ {
		for y := 0; y < tm.Height; y++ {
			smoothed[x][y] = tm.SmoothPoint(x, y, distance)
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

func (tm TileMap) RandomSmooth() [][]int {
	smoothed := make([][]int, tm.Width)
	for i := 0; i < tm.Width; i++ {
		smoothed[i] = make([]int, tm.Height)
	}
	for x := 0; x < tm.Width; x++ {
		for y := 0; y < tm.Height; y++ {
			smoothed[x][y] = tm.SmoothPoint(x, y, rand.Intn(3))
		}
	}
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
				Altitude: tm.Get(x, y),
				X:        x,
				Y:        y,
			}
			i++
		}
	}
	return tiles
}
