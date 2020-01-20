# FFmpeg go wrapper

## Example 1
```golang
NewFFmpeg().
    InputFile("testdata/input.mp4").
    OutputFile("testdata/output.mp4").
    AudioBitrate(126000).
    VideoBitrate(440000).
    Preset("fast").
    Seek(3).
    Duration(5).
    Run()
```

## Example 2

```golang
fout, _ := os.Create("testdata/output.mp4")
defer fout.Close()

fin, _ := os.Open("testdata/input.mp4")

NewFFmpeg().
    Input(fin).
    Output(fout).
    AudioBitrate(126000).
    VideoBitrate(440000).
    Preset("fast").
    Seek(3).
    Duration(5).
    Format("mp4").
    Movflags("+faststart").
    Run()
```