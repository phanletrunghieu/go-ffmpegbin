package ffmpegbin

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeVideoFile(t *testing.T) {
	ffmpeg := NewFFmpeg()
	err := ffmpeg.InputFile("testdata/input.mp4").
		OutputFile("testdata/output.mp4").
		AudioBitrate(126000).
		VideoBitrate(440000).
		Preset("fast").
		Seek(3).
		Duration(5).
		Run()
	assert.Nil(t, err)
}

func TestEncodeVideoBuffer(t *testing.T) {
	fout, err := os.Create("testdata/output.mp4")
	assert.Nil(t, err)
	defer fout.Close()

	fin, err := os.Open("testdata/input.mp4")
	assert.Nil(t, err)

	ffmpeg := NewFFmpeg()
	err = ffmpeg.Input(fin).
		Output(fout).
		AudioBitrate(126000).
		VideoBitrate(440000).
		Preset("fast").
		Seek(3).
		Duration(5).
		Format("mp4").
		Movflags("+faststart").
		RemoveMetadata(true).
		NoVideo(true).
		Run()
	assert.Nil(t, err)
}
