package audio

import (
	"fmt"

	"github.com/sandovalrr/mediacutter/models"
	"github.com/sandovalrr/mediacutter/utils"

	"github.com/coreos/pkg/capnslog"
)

var (
	log = capnslog.NewPackageLogger(models.Repo, "audio/cutter")
)

//Audio audio model
type Audio struct{}

//NewAudio return an instance of audio
func NewAudio() *Audio {
	return &Audio{}
}

//ProgramCommand return program name on windows and linux
func (audio *Audio) ProgramCommand() string {
	if utils.IsWindows() {
		return "sox.exe"
	}

	return "sox"
}

func (audio *Audio) lenProgramCommand() string {
	if utils.IsWindows() {
		return "soxi.exe"
	}

	return "soxi"
}

//LenCommand Return command to get media duration
func (audio *Audio) LenCommand(path string) string {
	return fmt.Sprintf("%s -D %s", audio.lenProgramCommand(), path)
}

//CutterCommand Return command to cut media in range
func (audio *Audio) CutterCommand(filePath, outputPath string, start, end int) string {

	return fmt.Sprintf("%s %s -b 16 -c 1 -r 16k %s trim %d =%d", audio.ProgramCommand(), filePath, outputPath, start, end)
}
