package comps

import (
	"fmt"
	"strconv"
)

type TileMap struct {
	MaxAltitude int
	Width       int
	Height      int
	SeedData    [][]int
	Tiles       [][]Tile
	Config      MapConfig
}

type Tile struct {
	Id       int
	Altitude int
	X        int
	Y        int
}

func (t Tile) Classes() string {
	fmt.Println("hght-" + strconv.Itoa(t.Altitude))
	return "hght-" + strconv.Itoa(t.Altitude)
}
