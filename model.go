package mediacutter

import "time"

//Repo repo
var Repo = "github.com/sandovalrr/mediacutter"

//CutterOption cutter options
type CutterOption struct {
	Name    string
	Samples time.Duration
}
