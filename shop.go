package smusic

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type shop struct {
	WeekdayStartHour   int
	WeekdayStartMinute int
	WeekdayStopHour    int
	WeekdayStopMinute  int
	WeekendStartHour   int
	WeekendStartMinute int
	WeekendStopHour    int
	WeekendStopMinute  int
	MediaFolders       []mediaFolder
}

func (s *shop) Configure() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("Config file not found")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(s)
	if err != nil {
		log.Fatalf("Config file corrupt.")
	}

	for i := range s.MediaFolders {
		if err = s.MediaFolders[i].loadMediaFiles(); err != nil {
			log.Fatalf("Can not load files from mediafolder %v.", s.MediaFolders[i].Path)
		}
	}
}

func (s *shop) StartHour() (hour int) {
	wday := time.Now().Weekday()
	if wday > 0 && wday < 6 {
		return s.WeekdayStartHour
	}
	return s.WeekendStartHour
}

func (s *shop) StartMinute() (minute int) {
	wday := time.Now().Weekday()
	if wday > 0 && wday < 6 {
		return s.WeekdayStartMinute
	}
	return s.WeekendStartMinute
}

func (s *shop) StopHour() (hour int) {
	wday := time.Now().Weekday()
	if wday > 0 && wday < 6 {
		return s.WeekdayStopHour
	}
	return s.WeekendStopHour
}

func (s *shop) StopMinute() (minute int) {
	wday := time.Now().Weekday()
	if wday > 0 && wday < 6 {
		return s.WeekdayStopMinute
	}
	return s.WeekendStopMinute
}
