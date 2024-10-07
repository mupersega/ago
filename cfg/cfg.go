package cfg

import "math/rand"

type InitialAltitudeModifier int

const (
	DeepWater InitialAltitudeModifier = -4
	Water     InitialAltitudeModifier = -2
	Land      InitialAltitudeModifier = 0
	Hill      InitialAltitudeModifier = 2
	Mountain  InitialAltitudeModifier = 4
)

type MapConfig struct {
	SelectiveDistance      int
	PostSmoothDistance     int
	WidthModifier          int
	InitialAltitude        InitialAltitudeModifier
	Mountains              int
	MountainAltitude       int
	MountainAltitudeWindow int
	MountainRadius         int
	MountainRadiusWindow   int
	MountainRanges         int
	MountainRangeSize      int
	RangeSpread            int
	DefaultRunners         int
	DefaultRunnerMinlength int
	DefaultRunnerMaxlength int
}

// RandomMountainRadius returns a random radius for the mountain based on the configuration values.
func (mc MapConfig) RandomMountainRadius() int {
	return mc.MountainRadius - rand.Intn(mc.MountainRadiusWindow*2+1) - mc.MountainRadiusWindow
}

func (mc MapConfig) RandomMountainAltitude() int {
	return mc.MountainAltitude - rand.Intn(mc.MountainAltitudeWindow*2+1) - mc.MountainAltitudeWindow
}
