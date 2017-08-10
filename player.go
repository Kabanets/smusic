package main

import (
	"io"

	"bitbucket.org/weberc2/media/ao"
	"bitbucket.org/weberc2/media/mpg123"
)

type player struct {
}

func (p *player) play(fileName string) error {

	mpg123.Initialize()
	defer mpg123.Exit()

	handle, err := mpg123.Open(fileName)
	if err == nil {
		defer handle.Close()

		ao.Initialize()
		defer ao.Shutdown()

		dev := ao.NewLiveDevice(aoSampleFormat(handle))
		defer dev.Close()

		_, err = io.Copy(dev, handle)
	}
	return err
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
