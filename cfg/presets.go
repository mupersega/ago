package cfg

func IslandsConfig() MapConfig {
	return MapConfig{
		SelectiveDistance:      5,
		WidthModifier:          1,
		PostSmoothDistance:     2,
		InitialAltitude:        DeepWater,
		Mountains:              0,
		MountainAltitude:       7,
		MountainAltitudeWindow: 0,
		MountainRadius:         10,
		MountainRadiusWindow:   2,
		MountainRanges:         2,
		MountainRangeSize:      4,
		RangeSpread:            30,
		DefaultRunners:         4,
		DefaultRunnerMinlength: 5,
		DefaultRunnerMaxlength: 10,
	}
}

func CanyonsConfig() MapConfig {
	return MapConfig{
		SelectiveDistance:      5,
		WidthModifier:          1,
		PostSmoothDistance:     2,
		InitialAltitude:        Mountain,
		Mountains:              0,
		MountainAltitude:       10,
		MountainAltitudeWindow: 4,
		MountainRadius:         5,
		MountainRadiusWindow:   4,
		MountainRanges:         10,
		MountainRangeSize:      10,
		RangeSpread:            30,
		DefaultRunners:         20,
		DefaultRunnerMinlength: 5,
		DefaultRunnerMaxlength: 10,
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
