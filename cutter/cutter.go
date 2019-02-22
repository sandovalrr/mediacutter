package cutter

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sandovalrr/mediacutter/audio"
	"github.com/sandovalrr/mediacutter/interfaces"
	"github.com/sandovalrr/mediacutter/models"
	"github.com/sandovalrr/mediacutter/utils"
	"github.com/sandovalrr/mediacutter/video"

	"github.com/coreos/pkg/capnslog"
)

var (
	log = capnslog.NewPackageLogger(models.Repo, "audio/cutter")
)

//Cutter cutter model
type Cutter struct {
	Options models.CutterOption
	Repo    interfaces.CutterRepository
}

//NewAudioCutter create a new instance of audio cutter
func NewAudioCutter(options models.CutterOption) *Cutter {
	repo := audio.NewAudio()
	return &Cutter{
		Options: options,
		Repo:    repo,
	}
}

//NewVideoCutter create a new instance of video cutter
func NewVideoCutter(options models.CutterOption) *Cutter {
	repo := video.NewVideo()
	return &Cutter{
		Options: options,
		Repo:    repo,
	}
}

//Len Returns media length in seconds
func (cutter *Cutter) Len() (time.Duration, error) {
	response := time.Duration(0)
	command := cutter.Repo.LenCommand(cutter.Options.Path)
	resp, err := utils.ExecCmd(command)
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

//Split split media file and returns an error if something wrong happen
func (cutter *Cutter) Split() error {

	ext := filepath.Ext(cutter.Options.Path)

	audioLen, err := cutter.Len()
	if err != nil {
		return err
	}
	totalChunks := int(audioLen / cutter.Options.Samples)
	mod := int(audioLen % cutter.Options.Samples)
	if mod > 0 {
		totalChunks++
	}

	chunkName := "%d-%d" + ext
	start := 0
	end := int(cutter.Options.Samples)
	if totalChunks == 1 {
		end = int(audioLen)
	}
	for i := 0; i < totalChunks; i++ {

		chunkPath := filepath.Join(cutter.Options.ChunkPath, fmt.Sprintf(chunkName, start, end))

		os.Remove(chunkPath)

		command := cutter.Repo.CutterCommand(cutter.Options.Path, chunkPath, start, end)
		_, err = utils.ExecCmd(command)

		if err != nil {
			log.Errorf("Receiving Error from sh => %v", err)
			continue
		}

		start += int(cutter.Options.Samples)
		end += int(cutter.Options.Samples)
		if i == totalChunks-2 {
			end -= mod
		}

	}

	return nil
}
