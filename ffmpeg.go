package ffmpegbin

import (
	"errors"
	"fmt"
	"io"

	"github.com/nickalie/go-binwrapper"
)

type FFmpeg struct {
	*binwrapper.BinWrapper

	inputFile string
	input     io.Reader

	outputFile string
	output     io.Writer

	format string

	audioBitrate int
	videoBitrate int

	videoCodec string
	audioCodec string

	seek     int
	duration int

	movflags string
	preset   string
}

func NewFFmpeg() *FFmpeg {
	bin := binwrapper.NewBinWrapper().ExecPath("ffmpeg").AutoExe().Debug()
	return &FFmpeg{
		BinWrapper: bin,
	}
}

func (f *FFmpeg) InputFile(file string) *FFmpeg {
	f.input = nil
	f.inputFile = file
	return f
}

func (f *FFmpeg) Input(reader io.Reader) *FFmpeg {
	f.inputFile = ""
	f.input = reader
	return f
}

func (f *FFmpeg) OutputFile(file string) *FFmpeg {
	f.output = nil
	f.outputFile = file
	return f
}

func (f *FFmpeg) Output(writer io.Writer) *FFmpeg {
	f.outputFile = ""
	f.output = writer
	return f
}

func (f *FFmpeg) Format(format string) *FFmpeg {
	f.format = format
	return f
}

func (f *FFmpeg) AudioBitrate(bitrate int) *FFmpeg {
	f.audioBitrate = bitrate
	return f
}

func (f *FFmpeg) VideoBitrate(bitrate int) *FFmpeg {
	f.videoBitrate = bitrate
	return f
}

func (f *FFmpeg) VideoCodec(codec string) *FFmpeg {
	f.videoCodec = codec
	return f
}

func (f *FFmpeg) AudioCodec(codec string) *FFmpeg {
	f.audioCodec = codec
	return f
}

func (f *FFmpeg) Seek(seek int) *FFmpeg {
	f.seek = seek
	return f
}

func (f *FFmpeg) Duration(duration int) *FFmpeg {
	f.duration = duration
	return f
}

func (f *FFmpeg) Movflags(movflags string) *FFmpeg {
	f.movflags = movflags
	return f
}

func (f *FFmpeg) Preset(preset string) *FFmpeg {
	f.preset = preset
	return f
}

// Ex: cat VID_20191112_093257.mp4 | ffmpeg -i - -ss 10 -t 6 -vcodec libx264 -acodec aac -b:a 10000 -b:v 10000 -movflags frag_keyframe+empty_moov+faststart -f mp4 pipe: > bbbbb.mp4
func (f *FFmpeg) Run() error {
	defer f.BinWrapper.Reset()

	// input
	if f.input != nil {
		f.Arg("-i", "-")
		f.StdIn(f.input)
	} else if f.inputFile != "" {
		f.Arg("-i", f.inputFile)
	} else {
		return errors.New("Undefined input")
	}

	// arg
	f.Arg("-y")
	if f.audioBitrate > 0 {
		f.Arg("-b:a", fmt.Sprintf("%d", f.audioBitrate))
	}
	if f.videoBitrate > 0 {
		f.Arg("-b:v", fmt.Sprintf("%d", f.videoBitrate))
	}
	if f.videoCodec != "" {
		f.Arg("-vcodec", f.videoCodec)
	}
	if f.audioCodec != "" {
		f.Arg("-acodec", f.audioCodec)
	}
	if f.seek >= 0 {
		f.Arg("-ss", fmt.Sprintf("%d", f.seek))
	}
	if f.duration > 0 {
		f.Arg("-t", fmt.Sprintf("%d", f.duration))
	}
	if f.format != "" {
		f.Arg("-f", f.format)
	}

	// output
	if f.outputFile != "" {
		f.Arg(f.outputFile)
	} else if f.output != nil {
		f.Arg("-movflags", f.movflags+"+frag_keyframe+empty_moov").Arg("pipe:")
		f.SetStdOut(f.output)
	} else {
		return errors.New("Undefined output")
	}

	// run
	err := f.BinWrapper.Run()
	if err != nil {
		return errors.New(err.Error() + ". " + string(f.StdErr()))
	}

	return nil
}
