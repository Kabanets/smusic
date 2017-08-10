package main

import (
	"io"

	"bitbucket.org/weberc2/media/ao"
	"bitbucket.org/weberc2/media/mpg123"
)

func main() {
	// Загрузить конфиг
	shop := new(shop)
	shop.configure()

	for i := 0; i <= 3; i++ {
		// По всем медиапапкам
		for _, mf := range shop.MediaFolders {
			// Определенное количество песен из папки за раз
			for i := 0; i < mf.Songs; i++ {
				f := mf.getNextMediaFile()

				println(f)

				mpg123.Initialize()
				defer mpg123.Exit()

				handle, err := mpg123.Open(mf.Path + "/" + f)
				if err != nil {
					continue
				}
				defer handle.Close()

				ao.Initialize()
				defer ao.Shutdown()
				dev := ao.NewLiveDevice(aoSampleFormat(handle))
				defer dev.Close()

				if _, err := io.Copy(dev, handle); err != nil {
					continue
				}
			}
		}
	}

}

// Get the ao.SampleFormat from the mpg123.Handle
func aoSampleFormat(handle *mpg123.Handle) *ao.SampleFormat {
	const bitsPerByte = 8

	rate, channels, encoding := handle.Format()

	return &ao.SampleFormat{
		BitsPerSample: handle.EncodingSize(encoding) * bitsPerByte,
		Rate:          int(rate),
		Channels:      channels,
		ByteFormat:    ao.FormatNative,
		Matrix:        nil,
	}
}
