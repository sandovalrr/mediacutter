package video

import (
	"fmt"

	"github.com/coreos/pkg/capnslog"
	"github.com/sandovalrr/mediacutter/models"
	"github.com/sandovalrr/mediacutter/utils"
)

var (
	log = capnslog.NewPackageLogger(models.Repo, "video/cutter")
)

//Video video model
type Video struct {
}

//NewVideo create an instance of Video
func NewVideo() *Video {
	return &Video{}
}

//ProgramCommand return program name on windows and linux
func (video *Video) ProgramCommand() string {
	if utils.IsWindows() {
		return "ffmpeg.exe"
	}

	return "ffmpeg"
}

func (video *Video) lenProgramCommand() string {
	if utils.IsWindows() {
		return "ffprobe.exe"
	}

	return "ffprobe"
}

//LenCommand Return command to get media duration
func (video *Video) LenCommand(path string) string {
	return fmt.Sprintf(`%s -i %s -show_entries format=duration -v quiet -of csv="p=0"`, video.lenProgramCommand(), path)
}

//CutterCommand Return command to cut media in range
func (video *Video) CutterCommand(filePath, outputPath string, start, end int) string {
	sTime := utils.ToTimeFormat(start)
	eTime := utils.ToTimeFormat(end)
	return fmt.Sprintf("%s -i %s -ss %s -t %s -async 1 -strict -2 %s", video.ProgramCommand(), filePath, sTime, eTime, outputPath)
}
