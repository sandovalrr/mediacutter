package models

import "time"

//Repo repo
var Repo = "github.com/sandovalrr/mediacutter"

//CutterOption cutter options
type CutterOption struct {
	Path      string
	Samples   time.Duration
	ChunkPath string
}
