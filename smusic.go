// Package smusic is a music player for shop with using Raspbbery PI
package smusic

import (
	"log"
	"time"
)

// Start function initialize media-folders and start to play music
func Start() {
	// Загрузить конфиг
	shop := new(shop)
	shop.Configure()

	for {
		// По всем медиапапкам
		for _, mf := range shop.MediaFolders {
			// Определенное количество песен из папки за раз
			for i := 0; i < mf.Songs; i++ {

				for h, m := getCurrentTime(); !((h > shop.StartHour() || (h == shop.StartHour() && m >= shop.StartMinute())) &&
					(h < shop.StopHour() || (h == shop.StopHour() && m < shop.StopMinute()))); h, m = getCurrentTime() {
					time.Sleep(60 * time.Second)
				}

				f := mf.getNextMediaFile()
				log.Println(f)

				p := new(player)
				if err := p.play(f); err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}
}

func getCurrentTime() (hour, minute int) {
	n := time.Now()
	return n.Hour(), n.Minute()
}
