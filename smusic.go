package main

import (
	"encoding/json"
	"os"
	// "time"
	// "bitbucket.org/weberc2/media/ao"
	// "bitbucket.org/weberc2/media/mpg123"
	"log"
)

type MediaFolder struct {
	Path  string
	Songs int
}

type Config struct {
	ShopId            string
	WeekdaysStartTime string
	WeekdaysStopTime  string
	WeekendsStartTime string
	WeekendsStopTime  string
	MainMediaPlan     []MediaFolder
}

// Get the ao.SampleFormat from the mpg123.Handle
// func aoSampleFormat(handle *mpg123.Handle) *ao.SampleFormat {
// 	const bitsPerByte = 8

// 	rate, channels, encoding := handle.Format()

// 	return &ao.SampleFormat{
// 		BitsPerSample: handle.EncodingSize(encoding) * bitsPerByte,
// 		Rate:          int(rate),
// 		Channels:      channels,
// 		ByteFormat:    ao.FormatNative,
// 		Matrix:        nil,
// 	}
// }

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Config file not found or corrupt.")
	}

	// if len(os.Args) < 2 {
	// 	os.Exit(0)
	// }

	// mpg123.Initialize()
	// defer mpg123.Exit()

	// handle, err := mpg123.Open("song.mp3")
	// if err != nil {
	// 	print(err.Error())
	// }
	// defer handle.Close()

	// ao.Initialize()
	// defer ao.Shutdown()
	// dev := ao.NewLiveDevice(aoSampleFormat(handle))
	// defer dev.Close()

	// if _, err := io.Copy(dev, handle); err != nil {
	// 	panic(err)
	// }
}
