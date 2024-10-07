package factory

import (
	"ago/vector"
	"encoding/json"
)

// this is the factory package which will build out all the shapes that we need
// so that the front end using three.js can render them

type Color struct {
	Hex         string  `json:"hex"`
	Transparent bool    `json:"transparent"`
	Opacity     float64 `json:"opacity"`
}

// Box represents a 3D box with width, height, depth, and position.
type Box struct {
	Width  float64     `json:"width"`
	Height float64     `json:"height"`
	Depth  float64     `json:"depth"`
	Pos    vector.Vec3 `json:"position"`
	Color  Color       `json:"color"`
}

// Boxes is a collection of Box objects.
type Boxes []Box

func GetColor(altitude int, opacity float64) Color {
	transparent := opacity < 1
	var hex string
	switch {
	case altitude == 0:
		hex = "#000040"
	case altitude == 1:
		hex = "#000080"
	case altitude == 2:
		hex = "#0000CD"
	case altitude == 3:
		hex = "#1E90FF"
	case altitude == 4:
		hex = "#2dc7ff"
	case altitude == 5:
		hex = "#02a50d"
	case altitude == 6:
		hex = "#1d6200"
	case altitude == 7:
		hex = "#6b4429"
	case altitude == 8:
		hex = "#838383"
	case altitude == 9:
		hex = "#EEE"
	case altitude == 10:
		hex = "#fff"
	default:
		hex = "#000"
	}
	return Color{hex, transparent, opacity}
}

func BoxFromTile(tile Tile) Box {
	return Box{
		Width:  1,
		Height: float64(tile.Altitude),
		Depth:  1,
		Color:  GetColor(tile.Altitude, 1),
		Pos: vector.Vec3{
			X: float64(tile.X),
			Y: float64(tile.Altitude) / 2,
			Z: float64(tile.Y)},
	}
}

func BoxesFromTileMap(tm TileMap) Boxes {
	var boxes Boxes
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			boxes = append(boxes, BoxFromTile(tm.Tiles[y][x]))
		}
	}
	return boxes
}

func WaterBoxFromTileMap(tm TileMap) Box {
	return Box{
		Width:  float64(tm.Width),
		Height: 4.8,
		Depth:  float64(tm.Height),
		Color:  GetColor(2, 0.5),
		Pos: vector.Vec3{
			X: float64(tm.Width)/2 - .5,
			Y: 2.4,
			Z: float64(tm.Height)/2 - .5,
		},
	}
}

// AsJson converts the Boxes collection to a JSON string.
func (b Boxes) AsJson() (string, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (b Box) AsJson() (string, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

type Sphere struct {
	Radius float64
}

type Cylinder struct {
	RadiusTop    float64
	RadiusBottom float64
	Height       float64
}

type Cone struct {
	Radius float64
	Height float64
}

// NewSphere creates a new Sphere with the specified radius.
func NewSphere(radius float64) Sphere {
	return Sphere{radius}
}

// NewCylinder creates a new Cylinder with the specified dimensions.
func NewCylinder(radiusTop, radiusBottom, height float64) Cylinder {
	return Cylinder{radiusTop, radiusBottom, height}
}

// NewCone creates a new Cone with the specified dimensions.
func NewCone(radius, height float64) Cone {
	return Cone{radius, height}
}
