package utils

import (
	"fmt"
	"math"
	"os/exec"
	"runtime"

	"github.com/coreos/pkg/capnslog"
	"github.com/sandovalrr/mediacutter/models"
)

var log = capnslog.NewPackageLogger(models.Repo, "utils/utils")

//IsWindows check if OS is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

//ExecCmd execute a command in terminal
func ExecCmd(cmd string, params ...interface{}) (string, error) {

	command := fmt.Sprintf(cmd, params...)
	log.Infof("Running Command => %s", command)

	var terminal string
	var path string

	if IsWindows() {
		terminal = "cmd.exe"
		path = "/C"
	} else {
		terminal = "sh"
		path = "-c"
	}

	data, err := exec.Command(terminal, path, command).CombinedOutput()

	var result string

	if data != nil && len(data) > 0 {
		result = string(data)
		log.Infof("Receiving from sh => %s", result)
	}

	return result, err
}

//ToTimeFormat return string in format time hh:mm:ss
func ToTimeFormat(duration int) string {
	hours := "00"
	minutes := "00"
	seconds := "00"

	if h := int(math.Floor(float64(duration) / 3600)); h > 0 {
		if h < 10 {
			hours = fmt.Sprintf("0%d", h)
		} else {
			hours = fmt.Sprintf("%d", h)
		}
	}

	if m := int(math.Floor(float64(duration)/60)) % 60; m > 0 {
		if m < 10 {
			minutes = fmt.Sprintf("0%d", m)
		} else {
			minutes = fmt.Sprintf("%d", m)
		}
	}

	if s := duration % 60; s > 0 {
		if s < 10 {
			seconds = fmt.Sprintf("0%d", s)
		} else {
			seconds = fmt.Sprintf("%d", s)
		}
	}

	return fmt.Sprintf("%s:%s:%s", hours, minutes, seconds)
}
