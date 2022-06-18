package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gbodra/swimming-calculator-api/model"
)

func RacePost(w http.ResponseWriter, r *http.Request) {
	var result map[string]interface{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	splits := result["splits"].([]interface{})
	var splitEntries []model.SplitEntry

	for _, split := range splits {
		v := split.(map[string]interface{})
		distance, _ := strconv.Atoi(fmt.Sprintf("%v", v["distance"]))
		hour, _ := strconv.Atoi(fmt.Sprintf("%v", v["hour"]))
		minute, _ := strconv.Atoi(fmt.Sprintf("%v", v["minute"]))
		second, _ := strconv.ParseFloat(fmt.Sprintf("%v", v["second"]), 32)

		splitEntry := model.NewSplitEntry(uint(distance), uint(hour), uint(minute), float32(second))
		splitEntries = append(splitEntries, splitEntry)

		log.Println("splitEntry:", splitEntry)
	}

	raceEntry := model.NewRaceEntry(splitEntries)
	log.Println("raceEntry:", raceEntry)

	w.Header().Set("Content-Type", "application/json")

	raceEntryJson, _ := json.Marshal(raceEntry)
	w.Write([]byte(raceEntryJson))
}
