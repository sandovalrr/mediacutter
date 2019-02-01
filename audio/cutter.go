package audio

import (
	"strconv"
	"strings"
	"time"

	"github.com/sandovalrr/mediacutter/utils"

	"github.com/coreos/pkg/capnslog"
	"github.com/sandovalrr/mediacutter"
)

var (
	log = capnslog.NewPackageLogger(mediacutter.Repo, "audio/cutter")
)

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

//Len Returns audio length in seconds
func (audio *Audio) Len() (time.Duration, error) {
	response := time.Duration(0)
	resp, err := utils.ExecCmd("soxi -D %s", audio.Options.Name)
	if err != nil {
		return response, err
	}
	value, err := strconv.ParseFloat(strings.TrimSpace(resp), 64)
	if err != nil {
		return response, err
	}

	response = time.Duration(int(value))

	return response, nil
}
