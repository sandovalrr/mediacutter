package audio

import (
	"fmt"
	"path/filepath"
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

func (audio *Audio) getSox() string {
	if utils.IsWindows() {
		return "sox.exe"
	}

	return "sox"
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

//Split split audio and returns an error if something wrong happen
func (audio *Audio) Split() error {
	audioLen, err := audio.Len()
	if err != nil {
		return err
	}
	totalChunks := int(audioLen / audio.Options.Samples)
	mod := int(audioLen % audio.Options.Samples)
	if mod > 0 {
		totalChunks++
	}

	command := "%s %s -b 16 -c 1 -r 16k %s trim %d =%d"
	chunkName := "%d-%d.wav"
	start := 0
	end := int(audio.Options.Samples)
	if totalChunks == 1 {
		end = int(audioLen)
	}
	for i := 0; i < totalChunks; i++ {

		chuckPath := filepath.Join(audio.Options.ChunkPath, fmt.Sprintf(chunkName, start, end))
		_, err := utils.ExecCmd(command, audio.getSox(), audio.Options.Name, chuckPath, start, end)

		if err != nil {
			log.Errorf("Receiving Error from sh => %v", err)
			continue
		}

		start += int(audio.Options.Samples)
		end += int(audio.Options.Samples)
		if i == totalChunks-2 {
			end -= mod
		}

	}

	return nil
}
