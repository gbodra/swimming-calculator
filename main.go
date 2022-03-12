package main

import (
	"log"

	"github.com/gbodra/swimming-calculator-api/model"
)

func main() {
	a := App{}
	a.Initialize()

	var splits []model.SplitEntry
	splitEntry1 := model.NewSplitEntry(50, 0, 0, 44.37)
	splits = append(splits, splitEntry1)

	splitEntry2 := model.NewSplitEntry(50, 0, 0, 48.18)
	splits = append(splits, splitEntry2)

	splitEntry3 := model.NewSplitEntry(50, 0, 0, 47.67)
	splits = append(splits, splitEntry3)

	splitEntry4 := model.NewSplitEntry(50, 0, 0, 48.08)
	splits = append(splits, splitEntry4)

	splitEntry5 := model.NewSplitEntry(50, 0, 0, 47.96)
	splits = append(splits, splitEntry5)

	splitEntry6 := model.NewSplitEntry(50, 0, 0, 48.25)
	splits = append(splits, splitEntry6)

	splitEntry7 := model.NewSplitEntry(50, 0, 0, 47.15)
	splits = append(splits, splitEntry7)

	splitEntry8 := model.NewSplitEntry(50, 0, 0, 51.47)
	splits = append(splits, splitEntry8)

	log.Println("Split entry: ", splits)

	raceEntry := model.NewRaceEntry(splits)

	log.Println("Race entry.TotalSeconds:", raceEntry.TotalSeconds)
	log.Println("Race entry.Pace:", raceEntry.Pace.ToString())
	log.Println("Race entry.Distance:", raceEntry.Distance)

	a.Run()
}
