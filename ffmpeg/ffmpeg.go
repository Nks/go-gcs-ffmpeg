package ffmpeg

import (
	"bytes"
	"github.com/Nks/go-gcs-ffmpeg/utils"
	"runtime"
	"strings"
)

type Configuration struct {
	FfmpegBin  string
	FfprobeBin string
}

func Configure() (Configuration, error) {
	var outFFmpeg bytes.Buffer
	var outProbe bytes.Buffer

	execFFmpegCommand := GetFFmpegExec()
	execFFprobeCommand := GetFFprobeExec()

	outFFmpeg, err := utils.TestCmd(execFFmpegCommand[0], execFFmpegCommand[1])
	if err != nil {
		return Configuration{}, err
	}

	outProbe, err = utils.TestCmd(execFFprobeCommand[0], execFFprobeCommand[1])
	if err != nil {
		return Configuration{}, err
	}

	ffmpeg := strings.Replace(outFFmpeg.String(), utils.LineSeparator(), "", -1)
	fprobe := strings.Replace(outProbe.String(), utils.LineSeparator(), "", -1)

	cnf := Configuration{ffmpeg, fprobe}
	return cnf, nil
}

func GetFFmpegExec() []string {
	var platform = runtime.GOOS
	var command = []string{"", "ffmpeg"}

	switch platform {
	case "windows":
		command[0] = "where"
		break
	default:
		command[0] = "which"
		break
	}

	return command
}

func GetFFprobeExec() []string {
	var platform = runtime.GOOS
	var command = []string{"", "ffprobe"}

	switch platform {
	case "windows":
		command[0] = "where"
		break
	default:
		command[0] = "which"
		break
	}
	return command
}
