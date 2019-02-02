<h1 align="center">Media Cutter</h1>

Simple Go Audio/Video cutter library using sox and ffmpeg tools.

## Instalation

```text
$ go get github.com/sandovalrr/mediacutter
```

or using glide

```text
$ glide get github.com/sandovalrr/mediacutter
```

## Usage

### Audio Cutter

```go
import (
  "github.com/sandovalrr/mediacutter/cutter"
  "github.com/sandovalrr/mediacutter/model"
)

//...
//...

audioCutter := cutter.NewAudioCutter(model.CutterOption{
  Path: "path_to_audio.mp3",
  Samples: 15,
  ChunkPath: "path_to_output_folder",
})

audioCutter.Split()

```

### Video Cutter

```go
import (
  "github.com/sandovalrr/mediacutter/cutter"
  "github.com/sandovalrr/mediacutter/model"
)

//...
//...

videoCutter := cutter.NewVideoCutter(model.CutterOption{
  Path: "path_to_video.avi",
  Samples: 15,
  ChunkPath: "path_to_output_folder",
})

videoCutter.Split()
```

## API

### CutterOption

| Property  | Description                                      | Type          |
| --------- | ------------------------------------------------ | ------------- |
| Path      | Path to source media file                        | string        |
| Samples   | Time duration in seconds for each splitted chunk | time.Duration |
| ChunkPath | Path to output chunk folder                      |

