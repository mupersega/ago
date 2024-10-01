package comps

import (
	"ago/vector"
	"fmt"
	"math"
	"math/rand"
)

type InitialAltitudeModifier int

func NewTileMap(maxAltitude, width, height int, config MapConfig) TileMap {
	tm := TileMap{
		MaxAltitude: maxAltitude,
		Width:       width,
		Height:      height,
		Config:      config,
	}
	// prep
	tm.SeedData = make([][]int, width)
	for i := 0; i < width; i++ {
		tm.SeedData[i] = make([]int, height)
	}
	tm.FromConfig()
	return tm
}

func (tm *TileMap) FromConfig() {
	tm.FillSeedData(int(tm.Config.InitialAltitude))
	fmt.Println(tm.Info())
	tm.SeedData = tm.SelectiveRandomSmooth(tm.Config.SelectiveDistance, tm.Width/tm.Config.WidthModifier)
	fmt.Println(tm.Info())
	tm.FormMountains()
	fmt.Println(tm.Info())
	tm.FormRanges()
	fmt.Println(tm.Info())
	tm.SeedData = tm.Smooth(tm.Config.PostSmoothDistance)
	tm.SeedData = tm.RandomSmooth(5)
	tm.Tiles = tm.GenerateTiles()
}

func (tm TileMap) Info() string {
	// return min and max vals
	return fmt.Sprintf("Width: %d, Height: %d, Max: %d, Min: %d, Class: %s", tm.Width, tm.Height, tm.MaxValue(), tm.MinValue(), tm.Class())
}

func (tm TileMap) MaxValue() int {
	max := 0
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			if tm.AltAt(x, y) > max {
				max = tm.AltAt(x, y)
			}
		}
	}
	return max
}

func (tm TileMap) MinValue() int {
	min := 10
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			if tm.AltAt(x, y) < min {
				min = tm.AltAt(x, y)
			}
		}
	}
	return min
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

func (tm *TileMap) FillSeedData(modifier int) {
	for x := 0; x < tm.Width; x++ {
		for y := 0; y < tm.Height; y++ {
			firstrand := rand.Intn(10) + 1
			secondrand := rand.Intn(10) + 1
			val := (firstrand + secondrand) / 2

			val += modifier

			// Ensure the value stays between 1 and 10
			if val < 1 {
				val = 1
			} else if val > 10 {
				val = 10
			}

			// Set the final value in the TileMap
			tm.Set(x, y, val)
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
			smoothed[y][x] = tm.SmoothPoint(vector.Vec2{
				X: float64(x),
				Y: float64(y),
			}, distance)
		}
	}
	return smoothed
}

// average the altitude of the surrounding tiles
func (tm TileMap) SmoothPoint(vec vector.Vec2, distance int) int {
	total := 0
	count := 0

	for dx := -distance; dx <= distance; dx++ {
		for dy := -distance; dy <= distance; dy++ {
			wrappedX := (int(vec.X) + dx + tm.Width) % tm.Width
			wrappedY := (int(vec.Y) + dy + tm.Height) % tm.Height

			total += tm.AltAt(wrappedX, wrappedY)
			count++
		}
	}
	trueAverage := float32(total) / float32(count)
	rounded := math.Round(float64(trueAverage))
	return int(rounded)
}

// smooth the points and their neighbours
func (tm TileMap) SmoothPointsAndNeighbours(points []vector.Vec2, distance int) [][]int {
	copied := make([][]int, tm.Width)
	for i, innerSlice := range tm.SeedData {
		copied[i] = make([]int, len(innerSlice))
		copy(copied[i], innerSlice)
	}
	// get neighbours and store
	neighbours := make([]vector.Vec2, 0)
	for _, point := range points {
		neighbours = append(neighbours, tm.GetNeighbours(point, distance)...)
	}
	for _, neighbour := range neighbours {
		copied[int(neighbour.Y)][int(neighbour.X)] = tm.SmoothPoint(vector.Vec2{
			X: neighbour.X,
			Y: neighbour.Y,
		}, distance)
	}
	return copied
}

func (tm TileMap) GetNeighbours(vec vector.Vec2, distance int) []vector.Vec2 {
	neighbours := make([]vector.Vec2, 0)
	for dx := -distance; dx <= distance; dx++ {
		for dy := -distance; dy <= distance; dy++ {
			wrappedX := (int(vec.X) + dx + tm.Width) % tm.Width
			wrappedY := (int(vec.Y) + dy + tm.Height) % tm.Height
			neighbours = append(neighbours, vector.Vec2{X: float64(wrappedX), Y: float64(wrappedY)})
		}
	}
	return neighbours
}

// smooth the entire map for a number of iterations using maxDistance to determine the cluster of tiles to smooth
func (tm TileMap) RandomSmooth(maxDistance int) [][]int {
	smoothed := make([][]int, tm.Width)
	for i := 0; i < tm.Width; i++ {
		smoothed[i] = make([]int, tm.Height)
	}
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			smoothed[y][x] = tm.SmoothPoint(vector.Vec2{
				X: float64(x),
				Y: float64(y),
			}, rand.Intn(maxDistance))
		}
	}
	return smoothed
}

// smooth pockets of tiles around the map for a number of iterations using maxDistance to determine the cluster of tiles to smooth
func (tm TileMap) SelectiveRandomSmooth(maxDistance int, numberOfPoints int) [][]int {
	smoothed := make([][]int, tm.Width)
	for i := 0; i < tm.Width; i++ {
		smoothed[i] = make([]int, tm.Height)
	}
	pointsToSmooth := make([]vector.Vec2, 0)
	for i := 0; i < numberOfPoints; i++ {
		x := rand.Intn(tm.Width)
		y := rand.Intn(tm.Height)
		pointsToSmooth = append(pointsToSmooth, vector.Vec2{
			X: float64(x),
			Y: float64(y),
		})
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

// FormMountains creates a number of mountains on the map.
func (tm TileMap) FormMountains() {
	for i := 0; i < tm.Config.Mountains; i++ {
		x := rand.Intn(tm.Width)
		y := rand.Intn(tm.Height)
		tm.CreateMountain(x, y, tm.Config.RandomMountainRadius(), tm.Config.RandomMountainAltitude(), tm.Config.DefaultRunners)
	}
}

func (tm TileMap) FormRanges() {
	for i := 0; i < tm.Config.MountainRanges; i++ {
		x := rand.Intn(tm.Width)
		y := rand.Intn(tm.Height)
		tm.CreateMountainRange(x, y)
	}
}

// RandomMountainRadius returns a random radius for the mountain based on the configuration values.
func (mc MapConfig) RandomMountainRadius() int {
	return mc.MountainRadius - rand.Intn(mc.MountainRadiusWindow*2+1) - mc.MountainRadiusWindow
}

func (mc MapConfig) RandomMountainAltitude() int {
	return mc.MountainAltitude - rand.Intn(mc.MountainAltitudeWindow*2+1) - mc.MountainAltitudeWindow
}

func (tm TileMap) CreateMountain(x, y, peakRadius, altitude int, runners int) {
	tm.DrawFilledCircle(x, y, peakRadius, altitude)
	// runners
	for i := 0; i < runners; i++ {
		// Random direction for the runner
		direction := vector.Vec2{X: rand.Float64()*2 - 1, Y: rand.Float64()*2 - 1}
		// Normalize the direction vector
		direction = direction.Normalize()
		// Create a runner from the center of the mountain
		tm.CreateRunner(vector.Vec2{X: float64(x), Y: float64(y)}, direction, altitude, tm.Config.DefaultRunnerMinlength, tm.Config.DefaultRunnerMaxlength)
	}
}

func (tm TileMap) CreateMountainRange(x, y int) {
	for i := 1; i <= tm.Config.MountainRangeSize; i++ {
		// Random radius for the mountain
		radius := tm.Config.RandomMountainRadius()
		// Random altitude for the mountain
		var altitude int
		if i == 1 {
			altitude = tm.Config.MountainAltitude
		} else {
			altitude = tm.Config.RandomMountainAltitude()
		}
		if altitude > 10 {
			fmt.Println("altitude > 10")
			fmt.Println(altitude)
		}
		// Random position for the mountain
		pos := vector.Vec2{X: float64(x + rand.Intn(tm.Config.RangeSpread) - 1), Y: float64(y + rand.Intn(tm.Config.RangeSpread) - 1)}
		tm.CreateMountain(int(pos.X), int(pos.Y), radius, altitude, tm.Config.DefaultRunners)
	}
}

func (tm TileMap) CreateRunner(start vector.Vec2, direction vector.Vec2, startAltitude int, minLength, maxLength int) {
	altitude := startAltitude
	currentPos := start

	// Loop to draw consecutive lines until the altitude reaches zero or we reach the map boundary
	for altitude > 0 {
		// Choose a random length for the next line within the range [minLength, maxLength]
		lineLength := minLength + rand.Intn(maxLength-minLength+1)

		// Calculate the end point for the line based on the random length
		endPos := currentPos.Add(direction.Normalize().Mul(float64(lineLength)))

		// Check if the end position is out of bounds
		if tm.isOutOfBounds(endPos) {
			break // stop drawing if we are at the edge of the map
		}

		// Draw the line between currentPos and endPos
		tm.DrawLine(currentPos, endPos, altitude)

		// Update the current position to the end of the line
		currentPos = endPos

		// Randomly vary the direction slightly to give natural movement
		rotationAngle := (rand.Float64() * math.Pi / 18) // random angle between 0 and 10°
		rotationDirection := rand.Intn(2)                // 0 = left, 1 = right
		if rotationDirection == 0 {
			rotationAngle = -rotationAngle // Rotate left (counter-clockwise)
		}
		direction = direction.Rotate(rotationAngle)

		// Decrease the altitude for the next line
		currentAltitude := tm.AltAt(int(currentPos.X), int(currentPos.Y))
		if currentAltitude >= altitude {
			break
		}
		altitude--

		// Chance to spawn a new runner from the current position
		if rand.Float64() < 0.6 { // 30% chance to spawn a new runner
			// More pronounced new direction (e.g., rotate by 45° to 90°)
			newDirectionAngle := math.Pi/4 + rand.Float64()*(math.Pi/4) // random between 45° and 90°
			newDirection := direction.Rotate(newDirectionAngle)

			// Spawn the new runner with the current position, new direction, and decreasing altitude
			go tm.CreateRunner(currentPos, newDirection, altitude, minLength, maxLength)
		}
	}
}

// DrawLine now takes an altitude as a parameter and draws a line with that altitude.
func (tm TileMap) DrawLine(startVec, endVec vector.Vec2, altitude int) {
	// Get the difference between the two vectors
	difference := endVec.Sub(startVec)
	// Get the magnitude of the difference
	magnitude := difference.Mag()
	// Get the normalized difference
	normalized := difference.Normalize()
	// Get the number of steps to take
	steps := int(magnitude)
	for i := 0; i < steps; i++ {
		// Get the new vector along the line
		newVec := startVec.Add(normalized.Mul(float64(i)))

		// Check if the new vector is out of bounds
		if tm.isOutOfBounds(newVec) {
			break // stop drawing if we are at the edge of the map
		}

		// Set the altitude for the current tile
		tm.Set(int(newVec.X), int(newVec.Y), altitude)
	}
}

// Helper function to check if a given position is out of the map bounds
func (tm TileMap) isOutOfBounds(pos vector.Vec2) bool {
	width, height := tm.Width, tm.Height // assuming TileMap has Width() and Height() methods
	return int(pos.X) < 0 || int(pos.X) >= width || int(pos.Y) < 0 || int(pos.Y) >= height
}

func (tm TileMap) DrawFilledRectangle(x, y, width, height, altitude int) {
	for dx := 0; dx < width; dx++ {
		for dy := 0; dy < height; dy++ {
			// Wrap coordinates using modulo
			wrappedX := (x + dx + tm.Width) % tm.Width
			wrappedY := (y + dy + tm.Height) % tm.Height
			// Set the new altitude
			tm.Set(wrappedX, wrappedY, altitude)
		}
	}
}

func (tm TileMap) DrawRectangle(x, y, width, height, altitude int) {
	for dx := 0; dx < width; dx++ {
		// Wrap coordinates using modulo
		wrappedX := (x + dx + tm.Width) % tm.Width
		// Set the new altitude
		tm.Set(wrappedX, y, altitude)
	}
	for dy := 0; dy < height; dy++ {
		// Wrap coordinates using modulo
		wrappedY := (y + dy + tm.Height) % tm.Height
		// Set the new altitude
		tm.Set(x, wrappedY, altitude)
	}
}

func (tm TileMap) DrawCircle(center vector.Vec2, radius float64, altitude int) {
	// get the number of steps to take
	steps := int(radius)
	for i := 0; i < steps; i++ {
		// get the angle
		angle := float64(i) * math.Pi / float64(steps)
		// get the new vector
		newVec := center.Add(vector.Vec2{X: math.Cos(angle) * radius, Y: math.Sin(angle) * radius})
		// set the altitude
		tm.Set(int(newVec.X), int(newVec.Y), altitude)
	}
}

func (tm TileMap) DrawFilledCircle(x, y, radius, altitude int) {
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			if dx*dx+dy*dy < radius*radius {
				// Wrap coordinates using modulo
				wrappedX := (x + dx + tm.Width) % tm.Width
				wrappedY := (y + dy + tm.Height) % tm.Height
				// Set the new altitude
				tm.Set(wrappedX, wrappedY, altitude)
			}
		}
	}
}
