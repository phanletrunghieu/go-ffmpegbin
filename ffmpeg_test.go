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
	fout, err := os.Create("testdata/output_buffer.mp4")
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

func TestExtractThumbnail(t *testing.T) {
	fout, err := os.Create("testdata/thumbnail.jpg")
	assert.Nil(t, err)
	defer fout.Close()

	ffmpeg := NewFFmpeg()
	err = ffmpeg.InputFile("testdata/input.mp4").
		Output(fout).
		Seek(1).
		Format("singlejpeg").
		VFrames(1).
		Run()
	assert.Nil(t, err)
}

func TestExtractGIFThumbnail(t *testing.T) {
	fout, err := os.Create("testdata/thumbnail.gif")
	assert.Nil(t, err)
	defer fout.Close()

	ffmpeg := NewFFmpeg()
	err = ffmpeg.InputFile("testdata/input.mp4").
		Output(fout).
		Seek(0).
		Duration(2).
		Rate(10).
		Loop(0).
		Format("gif").
		Run()
	assert.Nil(t, err)
}

func TestExtractGIFThumbnailWithScale(t *testing.T) {
	fout, err := os.Create("testdata/thumbnail-scale.gif")
	assert.Nil(t, err)
	defer fout.Close()

	ffmpeg := NewFFmpeg()
	err = ffmpeg.InputFile("testdata/input.mp4").
		Output(fout).
		Seek(0).
		Duration(2).
		Rate(10).
		Loop(0).
		FilterComplex("scale=800:-1").
		Format("gif").
		Run()
	assert.Nil(t, err)
}

func TestExtractTiktokWebpAnimated(t *testing.T) {
	fout, err := os.Create("testdata/thumbnail.webp")
	assert.Nil(t, err)
	defer fout.Close()

	ffmpeg := NewFFmpeg()
	err = ffmpeg.InputFile("testdata/input.mp4").
		Output(fout).
		Seek(2).
		Duration(2).
		Rate(30).
		CompressionLevel(4).
		QScale(70).
		FilterComplex("[0:v]scale=800:-1[vid];[0:v]scale=800:-1,reverse,fifo[r];[vid][r]concat,setpts=0.5*PTS [out]").
		Map("[out]").
		Format("webp").
		Run()
	assert.Nil(t, err)
}

func TestAddWatermarkToVideo(t *testing.T) {
	fout, err := os.Create("testdata/watermark.mp4")
	assert.Nil(t, err)
	defer fout.Close()

	ffmpeg := NewFFmpeg()
	err = ffmpeg.
		InputFile("testdata/input.mp4").
		InputFile("testdata/input.mp4").
		Output(fout).
		FilterComplex("[1:v]scale=77:-1 [ovr], [0:v][ovr] overlay=main_w-overlay_w*1.2:main_h-overlay_h*1.2").
		AudioCodec("copy").
		Format("mp4").
		Run()
	assert.Nil(t, err)
}
