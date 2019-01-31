package utils

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/coreos/pkg/capnslog"
	"github.com/sandovalrr/mediacutter"
)

var log = capnslog.NewPackageLogger(mediacutter.Repo, "utils/utils")

//IsWindows check if OS is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

//ExecCmd execute a command con terminal
func ExecCmd(cmd string, params ...interface{}) (string, error) {

	command := fmt.Sprintf(cmd, params...)
	log.Infof("Running Command => ", command)

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
		log.Infof("Receiving from sh => ", result)
	}

	return result, err
}
