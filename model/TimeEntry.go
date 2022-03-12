package model

import "fmt"

type SplitEntry struct {
	Distance     uint
	Hour         uint
	Minute       uint
	Second       float32
	TotalSeconds float32
}

func NewSplitEntry(distance uint, hour uint, min uint, sec float32) SplitEntry {
	entry := SplitEntry{
		Distance:     distance,
		Hour:         hour,
		Minute:       min,
		Second:       sec,
		TotalSeconds: getTotalSecondsSplit(hour, min, sec),
	}

	return entry
}

func getTotalSecondsSplit(hour uint, min uint, sec float32) float32 {
	totalSecs := float32(hour * 60 * 60)
	totalSecs += float32(min * 60)
	totalSecs += sec

	return totalSecs
}

type Pace struct {
	Minute uint
	Second uint
}

func (p Pace) ToString() string {
	minutes := p.Minute
	seconds := p.Second
	return fmt.Sprintf("%v:%v", minutes, seconds)
}

type RaceEntry struct {
	Splits       []SplitEntry
	TotalSeconds float32
	Pace         Pace
	Distance     uint
}

func NewRaceEntry(splits []SplitEntry) *RaceEntry {
	totalSecs := calculateRaceTotalSeconds(splits)
	distance := calculateRaceDistance(splits)

	entry := &RaceEntry{
		Splits:       splits,
		TotalSeconds: totalSecs,
		Pace:         calculateRacePace(totalSecs, distance),
		Distance:     distance,
	}

	return entry
}

func calculateRaceDistance(splits []SplitEntry) uint {
	var distance uint
	distance = 0
	for _, split := range splits {
		distance += split.Distance
	}

	return distance
}

func calculateRaceTotalSeconds(splits []SplitEntry) float32 {
	var totalSecs float32
	totalSecs = 0
	for _, split := range splits {
		totalSecs += split.TotalSeconds
	}

	return totalSecs
}

func calculateRacePace(totalSecs float32, distance uint) Pace {
	numberOf100m := distance / 100
	totalSecsPer100m := totalSecs / float32(numberOf100m)
	pace := Pace{
		Minute: uint(totalSecsPer100m) / 60,
		Second: uint(totalSecsPer100m) % 60,
	}

	return pace
}
