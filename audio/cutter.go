package audio

import (
	"github.com/coreos/pkg/capnslog"
	"github.com/sandovalrr/mediacutter"
)

var log = capnslog.NewPackageLogger(mediacutter.Repo, "audio/cutter")

//Audio audio model
type Audio struct {
	Options mediacutter.CutterOption
}

//NewAudioCutter create an instance of Audio
func NewAudioCutter(options mediacutter.CutterOption) *Audio {
	return &Audio{
		Options: options,
	}
}
