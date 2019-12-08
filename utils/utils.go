package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type TemporaryDirectory struct {
	path string
}

func DurToSec(dur string) (sec float64) {
	durAry := strings.Split(dur, ":")
	var secs float64
	if len(durAry) != 3 {
		return
	}
	hr, _ := strconv.ParseFloat(durAry[0], 64)
	secs = hr * (60 * 60)
	min, _ := strconv.ParseFloat(durAry[1], 64)
	secs += min * (60)
	second, _ := strconv.ParseFloat(durAry[2], 64)
	secs += second
	return secs
}

func LineSeparator() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

func TestCmd(command string, args string) (bytes.Buffer, error) {
	var out bytes.Buffer

	cmd := exec.Command(command, args)

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return out, err
	}

	return out, nil
}

func (tmpDir *TemporaryDirectory) CreateTempDirectory() (*TemporaryDirectory, error) {
	dir, err := ioutil.TempDir("", "gcs-transcoder")

	if err != nil {
		return nil, fmt.Errorf("unable create temporary directory %v", err)
	}

	return &TemporaryDirectory{path: dir}, nil
}

func (tmpDir *TemporaryDirectory) GetTempPath() string {
	return tmpDir.path
}

func (tmpDir *TemporaryDirectory) CleanUp() {
	if tmpDir.path != "" {
		err := os.RemoveAll(tmpDir.path)

		if err != nil {
			log.Println("unable delete temporary directory")
		}
	}
}
