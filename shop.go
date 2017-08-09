package main

import (
	"encoding/json"
	"log"
	"os"
)

type shop struct {
	ShopID           string
	WeekdayStartTime string
	WeekdayStopTime  string
	WeekendStartTime string
	WeekendStopTime  string
	MediaFolders     []mediaFolder
}

func (s *shop) configure() {
	file, err := os.Open("config.json")
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
