package comps

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

func IslandsConfig() MapConfig {
	return MapConfig{
		SelectiveDistance:      4,
		WidthModifier:          1,
		PostSmoothDistance:     2,
		InitialAltitude:        Water,
		Mountains:              0,
		MountainRanges:         2,
		MountainAltitude:       7,
		MountainAltitudeWindow: 4,
		MountainRadius:         5,
		MountainRadiusWindow:   1,
		MountainRangeSize:      3,
		RangeSpread:            20,
		DefaultRunners:         3,
		DefaultRunnerMinlength: 1,
		DefaultRunnerMaxlength: 5,
	}
}

func DefaultConfig() MapConfig {
	return MapConfig{
		SelectiveDistance:      5,
		WidthModifier:          1,
		PostSmoothDistance:     2,
		InitialAltitude:        Water,
		Mountains:              0,
		MountainAltitude:       10,
		MountainAltitudeWindow: 8,
		MountainRadius:         5,
		MountainRadiusWindow:   4,
		MountainRanges:         3,
		MountainRangeSize:      6,
		RangeSpread:            30,
		DefaultRunners:         10,
		DefaultRunnerMinlength: 5,
		DefaultRunnerMaxlength: 10,
	}
}
