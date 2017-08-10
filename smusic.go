package main

func main() {
	// Загрузить конфиг
	shop := new(shop)
	shop.configure()

	for {
		// По всем медиапапкам
		for _, mf := range shop.MediaFolders {
			// Определенное количество песен из папки за раз
			for i := 0; i < mf.Songs; i++ {
				f := mf.getNextMediaFile()
				println(f)
				p := new(player)
				if err := p.play(f); err != nil {
					println(err)
					continue
				}
			}
		}
	}
}
